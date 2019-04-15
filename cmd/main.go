package main

import (
	"flag"
	"fmt"
	"github.com/bregydoc/micro-culqi"
	"github.com/bregydoc/micro-culqi/chargers"
	"github.com/bregydoc/micro-culqi/messaging"
	"github.com/bregydoc/micro-culqi/proto"
	"github.com/bregydoc/micro-culqi/stores"
	"github.com/k0kubun/pp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	servicePort := flag.Int64("port", 18000, "Define the service port")
	configPath := flag.String("config", "/etc/uculqi/uculqi.config.yml", "Where your config file is it")
	flag.Parse()

	config, err := parseConfig(*configPath)
	if err != nil {
		panic(err)
	}

	pp.Println(config)

	charger := &chargers.CulqiCharger{
		ApiKey:       config.Culqi.ApiKey,
		MerchantCode: config.Culqi.MerchantCode,
	}
	mail := messaging.NewMailJetBackend(
		config.MailJet.User,
		config.MailJet.Password,
		config.MailJet.FromName,
		config.MailJet.FromEmail,
	)

	store, err := stores.NewBoltStore(config.Store.Filepath)
	if err != nil {
		panic(err)
	}

	if _, err := os.Stat(config.Company.TemplateFilename); os.IsNotExist(err) {
		err = ioutil.WriteFile(config.Company.TemplateFilename, []byte("Put your template here"), 0644)
		if err != nil {
			panic(err)
		}
	}

	s := &uculqi.Service{
		Company: &uculqi.Company{
			Info: &uculqi.CompanyInfo{
				Name:         config.Company.Name,
				PhoneSupport: config.Company.PhoneSupport,
				EmailSupport: config.Company.EmailSupport,
			},
			TemplateFilename: config.Company.TemplateFilename,
			Charger:          charger,
			Messaging:        mail,
			Store:            store,
			SendMaxAttempt:   config.Company.MaxAttempts,
		},
	}

	// Check if we have TLS a secure connection
	withTLS := true
	certificate := "/run/secrets/cert"
	key := "/run/secrets/key"

	_, err = os.Open(certificate)
	if err != nil || os.IsNotExist(err) {
		withTLS = false
	}
	if withTLS {
		_, err = os.Open(key)
		if err != nil || os.IsNotExist(err) {
			withTLS = false
		}
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *servicePort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var grpcServer *grpc.Server

	if withTLS {
		log.Println("setting with TLS from \"" + certificate + "\" and \"" + key + "\"")
		c, err := credentials.NewServerTLSFromFile(certificate, key)
		if err != nil {
			log.Fatalf("Failed to setup tls: %v", err)
		}
		grpcServer = grpc.NewServer(
			grpc.Creds(c),
		)
	} else {
		log.Println("setting without security")
		grpcServer = grpc.NewServer()
	}

	pculqi.RegisterUCulqiServer(grpcServer, s)

	log.Printf("[For Developers] GRPC listening on :%d\n", *servicePort)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}

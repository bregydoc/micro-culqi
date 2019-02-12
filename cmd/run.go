package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	"github.com/bregydoc/micro-culqi"
	"github.com/bregydoc/micro-culqi/grpc"
	"net"
)

func main() {
	servicePort := flag.Int64("port", 18000, "Define the service port")
	configPath := flag.String("config", "culqi.config.yml", "Where your config file is it")
	flag.Parse()

	conf, err := culqi.NewMCulqiConfig(*configPath)
	if err != nil {
		panic(err)
	}

	c := culqi.New(conf.MerchantCode, conf.APIKey, "v2", conf.ProductName)

	s := culqi.NewMCulqiService(c)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *servicePort))
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	mculqi.RegisterMCulqiServer(grpcServer, s)

	log.Printf("[For Developers] GRPC listening on :%d\n", *servicePort)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}

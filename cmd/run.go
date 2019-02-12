package main

import (
	"flag"

	"github.com/k0kubun/pp"

	culqi "github.com/bregydoc/micro-culqi"
)

func main() {
	// servicePort := flag.Int64("port", 18000, "Define the service port")
	configPath := flag.String("config", "culqi.config.yml", "Where your config file is it")
	flag.Parse()

	conf, err := culqi.NewMCulqiConfig(*configPath)
	if err != nil {
		panic(err)
	}

	c := culqi.New(conf.MerchantCode, conf.APIKey, "v2", conf.ProductName)

	s := newMCulqiService(c)
	pp.Println(s)
}

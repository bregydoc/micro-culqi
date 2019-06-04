package main

import (
	"log"

	"github.com/bregydoc/micro-culqi/bypass"
)

func main() {
	c := bypass.NewCulqi("<CULQI-API-KEY>")
	token, err := c.GenerateNewToken(&bypass.Details{
		Email:      "example@gexample.com",
		CardNumber: "4111111111111111",
		ExpYear:    2020,
		ExpMonth:   9,
		CVV:        "123",
	})
	if err != nil {
		panic(err)
	}

	log.Println(token)
}

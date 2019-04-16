package main

import (
	"context"
	"github.com/bregydoc/micro-culqi/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	creds, _ := credentials.NewClientTLSFromFile(
		"/Users/macbook/Documents/bombo/bombo-culqi/server.crt",
		"",
	)
	conn, err := grpc.Dial("localhost:18000", grpc.WithTransportCredentials(creds))
	// error handling omitted
	client := pculqi.NewUCulqiClient(conn)

	invoice, err := client.ExecuteCharge(context.TODO(), &pculqi.MinimalInvoice{
		Currency: pculqi.AvailableCurrency_PEN,
		Token:    "tkn_live_0CjjdWhFpEAZlxlz",
		Products: []*pculqi.Product{
			{Name: "Dry Herb Vaporizer", Currency: pculqi.AvailableCurrency_PEN, Price: 20.0},
		},
		Email: "bregymr@gmail.com",
	})
	if err != nil {
		panic(err)
	}

	log.Println("invoice code: " + invoice.Id + " was charged correctly")
}

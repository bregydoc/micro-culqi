package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bregydoc/micro-culqi/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	c, _ := credentials.NewClientTLSFromFile(
		"/Users/macbook/Documents/bombo/bombo-culqi/server.crt",
		"",
	)
	conn, err := grpc.Dial("localhost:18000", grpc.WithTransportCredentials(c))
	// error handling omitted
	client := pculqi.NewUCulqiClient(conn)

	// Creating the invoice
	invoice, err := client.CreateNewInvoice(context.TODO(), &pculqi.MinimalInvoice{
		Token:    "tkn_test_8thZUSDLMkozRQi7",
		Currency: pculqi.PEN,
		Products: []*pculqi.Product{
			{Name: "Dry Herb Vaporizer", Currency: pculqi.PEN, Price: 20.0},
		},
		Email: "customer@example.com",
	})
	if err != nil {
		panic(err)
	}

	// Charging the invoice
	invoice, err = client.ChargeInvoice(context.TODO(), &pculqi.InvoiceID{Id: invoice.Id})
	if err != nil {
		panic(err)
	}

	// Sending invoice as email
	invoice, err = client.SendInvoiceAsEmail(context.TODO(), &pculqi.InvoiceID{Id: invoice.Id})
	if err != nil {
		panic(err)
	}

	d, _ := json.Marshal(invoice)

	fmt.Println(string(d))
}

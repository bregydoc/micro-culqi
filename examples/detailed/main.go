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

	invoice, err := client.CreateNewInvoiceWithOrder(context.TODO(), &pculqi.Order{
		Token:    "tkn_test_8thZUSDLMkozRQi7",
		Currency: pculqi.PEN,
		Info: &pculqi.PersonInfo{
			Name:        "Bregy Malpartida",
			Email:       "bregymr@gmail.com",
			Phone:       957821858,
			CountryCode: "PE",
		},
		Card:     "****1111",
		Discount: 2.0,
		Products: []*pculqi.Product{
			{Name: "Dry Herb Vaporizer", Currency: pculqi.PEN, Price: 20.0},
		},
	})
	if err != nil {
		panic(err)
	}

	invoice, err = client.ChargeInvoice(context.TODO(), &pculqi.InvoiceID{Id: invoice.Id})
	if err != nil {
		panic(err)
	}

	d, _ := json.Marshal(invoice)

	fmt.Println(string(d))
}

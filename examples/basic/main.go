package main

import (
	"context"
	"log"

	pculqi "github.com/bregydoc/micro-culqi/proto"
	"github.com/k0kubun/pp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds, _ := credentials.NewClientTLSFromFile(
		"/Users/macbook/Documents/bombo/bombo-culqi/server.crt",
		"",
	)
	conn, err := grpc.Dial("localhost:18000", grpc.WithTransportCredentials(creds))
	// error handling omitted
	client := pculqi.NewUCulqiClient(conn)

	invoice, err := client.ExecuteChargeWithOrder(context.TODO(), &pculqi.Order{
		Id:       "AS1AG",
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
			{Name: "Dry Herb Vaporizer", Currency: pculqi.AvailableCurrency_PEN, Price: 20.0},
		},
	})

	if err != nil {
		panic(err)
	}

	log.Println("invoice code: " + invoice.Id + " was charged correctly")

	inv, err := client.GetInvoiceByID(context.TODO(), &pculqi.InvoiceID{
		Id: invoice.Id,
	})
	if err != nil {
		panic(err)
	}

	pp.Println(inv)
}

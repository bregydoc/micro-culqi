package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bregydoc/micro-culqi/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"time"
)

func main() {
	c, _ := credentials.NewClientTLSFromFile(
		"/Users/macbook/Documents/bombo/bombo-culqi/server.crt",
		"",
	)
	conn, err := grpc.Dial("localhost:18000", grpc.WithTransportCredentials(c))
	// error handling omitted
	client := pculqi.NewUCulqiClient(conn)

	t1 := time.Now()
	invoice, err := client.CreateNewInvoiceWithOrder(context.TODO(), &pculqi.Order{
		Token:    "tkn_test_zeMK1IsIS2OLhZNo",
		Currency: pculqi.PEN,
		Info: &pculqi.PersonInfo{
			Name:        "Bregy Malpartida",
			Email:       "bregymr@gmail.com",
			Phone:       957821858,
			CountryCode: "PE",
		},
		Card:     "****1111",
		Discount: 20.0,
		Products: []*pculqi.Product{
			{Name: "Dry Herb Vaporizer", Currency: pculqi.PEN, Price: 123.0},
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Invoice created, ID: %s in %s\n", invoice.Id, time.Since(t1))

	t1 = time.Now()
	invoice, err = client.ChargeInvoice(context.TODO(), &pculqi.InvoiceID{Id: invoice.Id})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Invoice charged, ID: %s in %s\n", invoice.Id, time.Since(t1))

	t1 = time.Now()
	invoice, err = client.SendInvoiceAsEmail(context.TODO(), &pculqi.InvoiceID{Id: invoice.Id})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Invoice sended, ID: %s in %s\n", invoice.Id, time.Since(t1))
	d, _ := json.Marshal(invoice)

	fmt.Println(string(d))
}

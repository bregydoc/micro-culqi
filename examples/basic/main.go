package main

import (
	"context"
	"github.com/bregydoc/micro-culqi/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:18000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pculqi.NewUCulqiClient(conn)

	invoice, err := client.ExecuteCharge(context.TODO(), &pculqi.MinimalInvoice{
		Currency: pculqi.AvailableCurrency_PEN,
		Token:    "<HERE YOUR CULQI TOKEN>",
		Products: []*pculqi.Product{
			{Name: "Dry Herb Vaporizer", Currency: pculqi.AvailableCurrency_PEN, Price: 20.0},
		},
		Email: "bregymr@gmail.com",
	})

	log.Println("invoice code: " + invoice.Id + " was charged correctly")
}

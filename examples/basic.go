package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bregydoc/micro-culqi/grpc"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:18000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	mCulqiClient := mculqi.NewMCulqiClient(conn)

	response, err := mCulqiClient.Charge(context.TODO(), &mculqi.ChargeParams{
		Currency: mculqi.Currency_PEN,
		Amount:   10,
		UserInfo: &mculqi.UserInformation{
			FirstName:   "Bregy",
			LastName:    "Malpartida",
			Email:       "bregymr@gmail.com",
			CountryCode: "PE",
			Token:       "<YOUR GENERATED TOKEN>",
		},
	})

	if err != nil {
		panic(err)
	}

	r, _ := json.Marshal(response)
	fmt.Println(string(r))
}

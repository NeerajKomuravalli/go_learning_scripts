package main

import (
	"context"
	"fmt"
	"log"

	"github.com/NeerajKomuravalli/go_learning_scripts/protobuf/celciusToFarenheit/proto/temperature"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", 9092), grpc.WithInsecure())
	if err != nil {
		log.Fatal("Unable to get a connection")
	}
	defer conn.Close()
	client := temperature.NewTempratureConversionClient(conn)

	// Convertsion from cal to far example
	resp, err := client.ConvertFromCelsiusToFarenheit(context.Background(), &temperature.Temprature{Temprature: 0})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Cal to Far => %v\n", resp)

	// Convertsion from far to cal example
	resp2, err := client.ConvertFromFarenheitToCelsius(context.Background(), &temperature.Temprature{Temprature: 0})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Far to Cal => %v\n", resp2)
}

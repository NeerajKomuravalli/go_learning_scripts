package main

import (
	"fmt"
	"log"
	"net"

	"github.com/NeerajKomuravalli/go_learning_scripts/protobuf/celciusToFarenheit/proto/temperature"
	"github.com/NeerajKomuravalli/go_learning_scripts/protobuf/celciusToFarenheit/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	gs := grpc.NewServer()
	server := server.NewServer()

	// register the temepfrature conversion server
	temperature.RegisterTempratureConversionServer(gs, server)

	// register the reflection service which allows clients to determine the methods
	// for this gRPC service
	reflection.Register(gs)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", 9092))
	if err != nil {
		log.Fatal("Unable to listen")
	}
	gs.Serve(listen)
}

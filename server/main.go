package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"server/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
)

func main() {
	var s *grpc.Server

	tls := flag.Bool("tls", false, "use a secure TLS connection")
	flag.Parse() //so we can use flag
	// flag.Bool -> tls as pointer of Bool
	if *tls {
		//cer file and key file
		certFile := "../tls/server.crt"
		keyFile := "../tls/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatal(err)
		}
		s = grpc.NewServer(grpc.Creds(creds))
	} else {
		//unsecure
		s = grpc.NewServer()
	}
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	services.RegisterCalculatorServer(s, services.NewCalculatorServer())
	//mine reflection not working idk why, grpcurl -v -plaintext -d '{"name": "k"}' localhost:50051 services.Calculator.Hello grpcurl working tho
	reflection.Register(s) //for evans, so you don't have to define its proto everytime
	//create reflection after the server is created and before the server starts serving requests.

	fmt.Print("gRPC server listening on port 50051")
	if *tls {
		fmt.Println(" with TLS")
	} else {
		fmt.Println()
	}
	err = s.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

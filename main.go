package main

import (
	"crypto/tls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	pb "server/grpc"
)

const (
	address  = ":5001"
	certFile = "cert.crt"
	keyFile  = "cert.key"
)

var games []*Game

func main() {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatalf("Error loading ssl certificate {%s}\n", err)
	}

	config := tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	listener, err := tls.Listen("tcp", address, &config)

	if err != nil {
		log.Fatalf("Error starting server {%s}\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterGameServiceServer(s, &GameService{})
	log.Println("Listening on:", listener.Addr())

	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalln("Failed to server:", err)
	}
}

package main

import (
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	pb "server/grpc"
)

const (
	address  = ":5001"
	certFile = "cert.crt"
	keyFile  = "cert.key"
)

var games map[uuid.UUID]*Game
var players map[string]*Player

func main() {
	games = make(map[uuid.UUID]*Game)
	players = make(map[string]*Player)
	/*
		cert, err := tls.LoadX509KeyPair(certFile, keyFile)
		if err != nil {
			log.Fatalf("Error loading ssl certificate {%s}\nAdd cert.crt and cert.key files\n", err)
		}

		config := tls.Config{
			Certificates: []tls.Certificate{cert},
		}

		listener, err := tls.Listen("tcp", address, &config)
	*/

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error starting server {%s}\n", err)
	}

	s := grpc.NewServer()

	pb.RegisterGameServiceServer(s, &GameService{})
	pb.RegisterPlayerServiceServer(s, &PlayerService{})
	pb.RegisterCardServiceServer(s, &CardService{})

	log.Println("Listening on:", listener.Addr())

	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalln("Failed to server:", err)
	}
}

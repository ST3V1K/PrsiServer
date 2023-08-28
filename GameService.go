package main

import (
	"context"
	"google.golang.org/grpc/peer"
	"log"
	pb "server/grpc"
)

type GameService struct {
	pb.GameServiceServer
}

func (s *GameService) Join(ctx context.Context, in *pb.GameJoinRequest) (*pb.GameJoinResponse, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		log.Fatalln("Cannot get peer from context")
	}
	log.Println("Received join request from:", p.Addr)

	return &pb.GameJoinResponse{
		Joined: true,
	}, nil
}

func (s *GameService) NewGame(ctx context.Context, in *pb.GameCreateRequest) (*pb.GameCreateResponse, error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		log.Fatalln("Cannot get peer from context")
	}
	player := &Player{
		Player:    in.Player,
		IPAddress: p.Addr,
	}
	game := NewGame(player)

	return &pb.GameCreateResponse{
		Uuid: game.Uuid,
	}, nil
}

func (s *GameService) ListGames(ctx context.Context, in *pb.ListGamesRequest) (*pb.ListGamesResponse, error) {
	var ids []*pb.Game

	for _, game := range games {
		ids = append(ids, &pb.Game{Uuid: game.Uuid})
	}

	return &pb.ListGamesResponse{
		Games: ids,
	}, nil
}

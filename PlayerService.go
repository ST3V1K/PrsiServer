package main

import (
	"context"
	"log"
	pb "server/grpc"
)

type PlayerService struct {
	pb.PlayerServiceServer
}

func (s *PlayerService) Login(ctx context.Context, in *pb.PlayerName) (*pb.PlayerPassword, error) {
	if _, exists := players[in.Name]; exists {
		return &pb.PlayerPassword{}, nil
	}

	player := NewPlayer(in.Name)
	log.Println(in.Name, "logged in")

	players[in.Name] = player
	return &pb.PlayerPassword{
		Password: player.Password,
	}, nil
}

func (s *PlayerService) Logout(ctx context.Context, in *pb.GameRequest) (*pb.SuccessResponse, error) {
	if !ValidatePlayer(in.Player) {
		return &pb.SuccessResponse{}, nil
	}

	log.Println(in.Player.Name, "logged out")

	Surrender(in.Game, in.Player)
	delete(players, in.Player.Name)

	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

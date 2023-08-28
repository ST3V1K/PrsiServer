package main

import (
	"context"
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

	players[in.Name] = player
	return &pb.PlayerPassword{
		Password: player.Password,
	}, nil
}

func (s *PlayerService) Logout(ctx context.Context, in *pb.Player) (*pb.SuccessResponse, error) {
	// TODO: remove player from active game and notify other of winning

	if !ValidatePlayer(in) {
		return &pb.SuccessResponse{}, nil
	}

	delete(players, GetPlayer(in).Name)
	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

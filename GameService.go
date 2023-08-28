package main

import (
	"context"
	pb "server/grpc"
)

type GameService struct {
	pb.GameServiceServer
}

func (s *GameService) Join(ctx context.Context, in *pb.GameRequest) (*pb.GameJoinResponse, error) {
	if !ValidatePlayer(in.Player) {
		return &pb.GameJoinResponse{}, nil
	}

	player := GetPlayer(in.Player)

	if player.InGame {
		return &pb.GameJoinResponse{}, nil
	}

	game := GetGameFromStringId(in.Game.Uuid)
	var opponentName string
	for opponentName = range game.Players {
		break
	}
	game.Players[player.Name] = player

	return &pb.GameJoinResponse{
		OpponentName: opponentName,
	}, nil
}

func (s *GameService) Leave(ctx context.Context, in *pb.GameRequest) (*pb.SuccessResponse, error) {
	//TODO: Inform player of winning

	if !ValidatePlayer(in.Player) {
		return &pb.SuccessResponse{}, nil
	}

	game := GetGameFromStringId(in.Game.Uuid)
	game.DisconnectPlayer(in.Player.Name)

	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

func (s *GameService) NewGame(ctx context.Context, in *pb.Player) (*pb.GameCreateResponse, error) {
	if !ValidatePlayer(in) {
		return &pb.GameCreateResponse{}, nil
	}

	player := GetPlayer(in)

	if player.InGame {
		return &pb.GameCreateResponse{}, nil
	}

	game := NewGame(player)
	return &pb.GameCreateResponse{
		Uuid: game.Uuid,
	}, nil
}

func (s *GameService) ListGames(ctx context.Context, in *pb.Player) (*pb.ListGamesResponse, error) {
	if !ValidatePlayer(in) {
		return &pb.ListGamesResponse{}, nil
	}

	var ids []*pb.Game

	for id := range games {
		ids = append(ids, &pb.Game{
			Uuid: id.String(),
		})
	}

	return &pb.ListGamesResponse{
		Games: ids,
	}, nil
}

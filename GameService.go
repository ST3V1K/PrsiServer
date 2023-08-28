package main

import (
	"context"
	"google.golang.org/grpc/peer"
	pb "server/grpc"
)

type GameService struct {
	pb.GameServiceServer
}

func (s *GameService) Join(ctx context.Context, in *pb.GameRequest) (*pb.GameJoinResponse, error) {
	p, ok := peer.FromContext(ctx)
	player := GetOrCreatePlayer(in.Player, p.Addr)
	game := GetGameFromStringId(in.Game.Uuid)

	if !ok || player.InGame {
		return &pb.GameJoinResponse{}, nil
	}

	game.Players = append(game.Players, player)
	return &pb.GameJoinResponse{
		OpponentName: game.Players[0].Name,
	}, nil
}

func (s *GameService) Leave(ctx context.Context, in *pb.GameRequest) (*pb.SuccessResponse, error) {
	//TODO: Inform player of winning

	game := GetGameFromStringId(in.Game.Uuid)

	for i, player := range game.Players {
		if player.Equals(in.Player) {
			game.Players = append(game.Players[:i], game.Players[i+1:]...)
			return &pb.SuccessResponse{
				Success: true,
			}, nil
		}
	}

	return &pb.SuccessResponse{
		Success: false,
	}, nil
}

func (s *GameService) NewGame(ctx context.Context, in *pb.PlayerRequest) (*pb.GameCreateResponse, error) {
	p, ok := peer.FromContext(ctx)
	player := GetOrCreatePlayer(in.Player, p.Addr)

	if !ok || player.InGame {
		return &pb.GameCreateResponse{}, nil
	}

	game := NewGame(player)
	return &pb.GameCreateResponse{
		Uuid: game.Uuid,
	}, nil
}

func (s *GameService) ListGames(ctx context.Context, in *pb.PlayerRequest) (*pb.ListGamesResponse, error) {
	var ids []*pb.Game

	for id := range games {
		ids = append(ids, &pb.Game{Uuid: id.String()})
	}

	return &pb.ListGamesResponse{
		Games: ids,
	}, nil
}

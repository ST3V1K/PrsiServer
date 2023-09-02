package main

import (
	"context"
	pb "server/grpc"
)

type CardService struct {
	pb.CardServiceServer
}

func (s *CardService) Play(ctx context.Context, in *pb.PlayCardRequest) (*pb.SuccessResponse, error) {
	if !ValidatePlayer(in.Player) {
		return &pb.SuccessResponse{}, nil
	}

	game := GetGameFromStringId(in.Game.Uuid)
	game.PlayCard(in.Card, in.Player.Name)

	for name, player := range game.Players {
		if name != in.Player.Name {
			err := player.InformPlayerOfPlay(&pb.GameStream{
				Uuid: game.Uuid,
				Seed: game.Seed,
			}, in.Card)
			if err != nil {
				panic(err)
				return nil, err
			}
			break
		}
	}

	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

func (s *CardService) Draw(ctx context.Context, in *pb.DrawCardRequest) (*pb.SuccessResponse, error) {
	if !ValidatePlayer(in.Player) {
		return &pb.SuccessResponse{}, nil
	}

	game := GetGameFromStringId(in.Game.Uuid)
	game.DrawCard(in.Draw, in.Player.Name)

	for name, player := range game.Players {
		if name != in.Player.Name {
			err := player.InformPlayerOfDraw(&pb.GameStream{
				Uuid: game.Uuid,
				Seed: game.Seed,
			}, &in.Draw)
			if err != nil {
				panic(err)
				return nil, err
			}
			break
		}
	}

	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

func (s *CardService) Stand(ctx context.Context, in *pb.StandRequest) (*pb.SuccessResponse, error) {
	if !ValidatePlayer(in.Player) {
		return &pb.SuccessResponse{}, nil
	}

	game := GetGameFromStringId(in.Game.Uuid)

	for name, player := range game.Players {
		if name != in.Player.Name {
			err := player.InformPlayerOfStand()
			if err != nil {
				panic(err)
				return nil, err
			}
			break
		}
	}

	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

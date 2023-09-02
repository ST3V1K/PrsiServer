package main

import (
	"context"
	"github.com/google/uuid"
	pb "server/grpc"
	"strings"
)

type GameService struct {
	pb.GameServiceServer
}

func (s *GameService) Join(in *pb.GameRequest, stream pb.GameService_JoinServer) error {
	if !ValidatePlayer(in.Player) {
		return nil
	}

	player := GetPlayer(in.Player)

	if player.InGame {
		return nil
	}

	game := GetGameFromStringId(in.Game.Uuid)
	gameStream := &pb.GameStream{
		Uuid: game.Uuid,
		Seed: game.Seed,
	}

	for opponentName := range game.Players {
		gameStream.OpponentName = &opponentName
		break
	}
	game.Players[player.Name] = player

	if err := stream.Send(gameStream); err != nil {
		return err
	}

	player.StreamNew = nil
	player.StreamJoin = &stream
	for !game.Finished {
		//TODO
	}

	return nil
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

func (s *GameService) NewGame(in *pb.Player, stream pb.GameService_NewGameServer) error {
	if !ValidatePlayer(in) {
		return nil
	}

	player := GetPlayer(in)

	if player.InGame {
		return nil
	}

	game := NewGame(player)
	gameStream := &pb.GameStream{
		Uuid: game.Uuid,
		Seed: game.Seed,
	}

	err := stream.Send(gameStream)

	if err != nil {
		panic(err)
		return err
	}

	for len(game.Players) <= 1 {
	}

	var i int
	for name := range game.Players {
		if i == len(game.Players)-1 {
			gameStream.OpponentName = &name
			break
		}
		i++
	}

	err = stream.Send(gameStream)
	if err != nil {
		panic(err)
		return err
	}

	player.StreamNew = &stream
	player.StreamJoin = nil
	for !game.Finished {
		//TODO
	}

	return nil
}

func (s *GameService) RemoveGame(ctx context.Context, in *pb.GameRequest) (*pb.SuccessResponse, error) {
	if !ValidatePlayer(in.Player) {
		return &pb.SuccessResponse{}, nil
	}

	id := uuid.MustParse(in.Game.Uuid)

	if _, ok := games[id]; !ok {
		return &pb.SuccessResponse{}, nil
	}

	delete(games, id)
	return &pb.SuccessResponse{
		Success: true,
	}, nil
}

func (s *GameService) ListGames(ctx context.Context, in *pb.Player) (*pb.ListGamesResponse, error) {
	if !ValidatePlayer(in) {
		return &pb.ListGamesResponse{}, nil
	}

	var ids []string
	var names []string

	for id, game := range games {
		ids = append(ids, id.String())
		for _, player := range game.Players {
			names = append(names, player.Name)
		}
	}

	return &pb.ListGamesResponse{
		GameUuids:   ids,
		PlayerNames: names,
	}, nil
}

func (s *GameService) ListGamesFiltered(ctx context.Context, in *pb.FilteredGamesRequest) (*pb.ListGamesResponse, error) {
	if !ValidatePlayer(in.Player) {
		return &pb.ListGamesResponse{}, nil
	}

	var ids []string
	var names []string

	for id, game := range games {
		var name string
		for _, player := range game.Players {
			name = player.Name
			break
		}

		if !strings.Contains(name, in.Filter) {
			continue
		}

		ids = append(ids, id.String())
		names = append(names, name)
	}

	return &pb.ListGamesResponse{
		GameUuids:   ids,
		PlayerNames: names,
	}, nil
}

func (s *GameService) RequestTie(ctx context.Context, in *pb.GameRequest) (*pb.SuccessResponse, error) {
	if !ValidatePlayer(in.Player) {
		return &pb.SuccessResponse{}, nil
	}

	game := GetGameFromStringId(in.Game.Uuid)

	for name, player := range game.Players {
		if name != in.Player.Name {
			err := player.InformPlayerOfTie()
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

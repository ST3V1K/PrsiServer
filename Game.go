package main

import (
	"github.com/google/uuid"
	"math/rand"
	pb "server/grpc"
)

type Game struct {
	*pb.Game
	Players []*Player
	Seed    int `default:"rand.Int()"`
}

func NewGame(player *Player) *Game {
	game := &Game{
		Game: &pb.Game{
			Uuid: uuid.New().String(),
		},
		Players: []*Player{player},
		Seed:    rand.Int(),
	}
	games = append(games, game)

	return game
}

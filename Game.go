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
	_uuid := uuid.New()

	game := &Game{
		Game: &pb.Game{
			Uuid: _uuid.String(),
		},
		Players: []*Player{player},
		Seed:    rand.Int(),
	}

	player.InGame = true
	players = append(players, player)
	games[_uuid] = game
	return game
}

func GetGameFromStringId(id string) *Game {
	return games[uuid.MustParse(id)]
}

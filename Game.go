package main

import (
	"github.com/google/uuid"
	"math/rand"
	pb "server/grpc"
)

type Game struct {
	*pb.Game
	Players map[string]*Player
	Seed    int
}

func NewGame(player *Player) *Game {
	_uuid := uuid.New()

	playersMap := make(map[string]*Player)
	playersMap[player.Name] = player

	game := &Game{
		Game: &pb.Game{
			Uuid: _uuid.String(),
		},
		Players: playersMap,
		Seed:    rand.Int(),
	}

	player.InGame = true
	players[player.Name] = player
	games[_uuid] = game
	return game
}

func GetGameFromStringId(id string) *Game {
	return games[uuid.MustParse(id)]
}

func (game *Game) DisconnectPlayer(playerName string) {
	delete(game.Players, playerName)
	if len(game.Players) == 0 {
		delete(games, uuid.MustParse(game.Uuid))
	}
}

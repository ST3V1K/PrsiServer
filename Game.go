package main

import (
	"github.com/google/uuid"
	"math/rand"
	pb "server/grpc"
)

type Game struct {
	*pb.Game
	Players  map[string]*Player
	Seed     int32
	LastCard *pb.Card
	Finished bool
}

func NewGame(player *Player) *Game {
	_uuid := uuid.New()

	//player.FillHand()

	playersMap := make(map[string]*Player)
	playersMap[player.Name] = player

	game := &Game{
		Game: &pb.Game{
			Uuid: _uuid.String(),
		},
		Players: playersMap,
		Seed:    rand.Int31(),
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

func (game *Game) PlayCard(card *pb.Card, playerName string) {
	for name, player := range game.Players {
		if name != playerName {
			if _, ok := player.Hand[card.String()]; ok {
				delete(player.Hand, card.String())
			}
			break
		}
	}

	game.LastCard = card
}

func (game *Game) DrawCard(draw int32, playerName string) {
	for name := range game.Players {
		if name != playerName {
			// TODO: add to hand
			break
		}
	}
}

func (game *Game) CanBePlayed(card *pb.Card) bool {
	// TODO: be more specific
	/*
		color := game.LastCard.Color == card.Color
		number := game.LastCard.Value == card.Value
		changeColor := card.Value == 12
		wasChanged := game.LastCard.Value == 12

		return color || number || changeColor || wasChanged
	*/
	return true
}

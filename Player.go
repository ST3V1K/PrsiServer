package main

import (
	"math/rand"
	pb "server/grpc"

	"github.com/google/uuid"
)

type Player struct {
	*pb.Player
	InGame     bool
	Hand       map[string]*pb.Card
	StreamNew  *pb.GameService_NewGameServer
	StreamJoin *pb.GameService_JoinServer
}

func NewPlayer(name string) *Player {
	return &Player{
		Player: &pb.Player{
			Name:     name,
			Password: RandomPassword(),
		},
	}
}

func GetPlayer(player *pb.Player) *Player {
	return players[player.Name]
}

func RandomPassword() string {
	const passLen = 16
	const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_.@#$%&*"

	b := make([]byte, passLen)
	for i := range b {
		b[i] = symbols[rand.Int63()%int64(len(symbols))]
	}
	return string(b)
}

func ValidatePlayer(player *pb.Player) bool {
	if player == nil {
		return false
	}
	if p, ok := players[player.Name]; ok {
		return p.Password == player.Password
	}
	return false
}

func Surrender(game *pb.Game, player *pb.Player) {
	// TODO
	id, _ := uuid.Parse(game.GetUuid())

	if game, ok := games[id]; ok {
		if len(game.Players) == 0 {
			delete(games, id)
		}
	}
}

func (p *Player) FillHand(seed int) {
	//TODO
}

func (p *Player) InformPlayerOfPlay(gameStream *pb.GameStream, card *pb.Card) error {
	gameStream.Card = card
	if p.StreamJoin != nil {
		return (*p.StreamJoin).Send(gameStream)
	}
	return (*p.StreamNew).Send(gameStream)
}

func (p *Player) InformPlayerOfDraw(gameStream *pb.GameStream, draw *int32) error {
	gameStream.Draw = draw
	if p.StreamJoin != nil {
		return (*p.StreamJoin).Send(gameStream)
	}
	return (*p.StreamNew).Send(gameStream)
}

func (p *Player) InformPlayerOfStand() error {
	if p.StreamJoin != nil {
		return (*p.StreamJoin).Send(&pb.GameStream{})
	}
	return (*p.StreamNew).Send(&pb.GameStream{})
}

func (p *Player) InformPlayerOfTie() error {
	out := true
	if p.StreamJoin != nil {
		return (*p.StreamJoin).Send(&pb.GameStream{Tie: &out})
	}
	return (*p.StreamNew).Send(&pb.GameStream{Tie: &out})
}

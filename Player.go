package main

import (
	"math/rand"
	pb "server/grpc"
)

type Player struct {
	*pb.Player
	InGame bool
	Hand   []any
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

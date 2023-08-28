package main

import (
	"net"
	pb "server/grpc"
)

type Player struct {
	*pb.Player
	IPAddress net.Addr
	InGame    bool
	Hand      []any
}

func (player *Player) Equals(player2 *pb.Player) bool {
	return player.Name == player2.Name &&
		player.Password == player2.Password
}

func GetOrCreatePlayer(player *pb.Player, ip net.Addr) *Player {
	for _, p := range players {
		if p.Equals(player) {
			return p
		}
	}
	return &Player{
		Player:    player,
		IPAddress: ip,
	}
}

func GetPlayer(player *pb.Player) *Player {
	for _, p := range players {
		if p.Equals(player) {
			return p
		}
	}
	return nil
}

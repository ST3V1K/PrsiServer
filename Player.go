package main

import (
	"net"
	pb "server/grpc"
)

type Player struct {
	*pb.Player
	IPAddress net.Addr
	Hand      []any
}

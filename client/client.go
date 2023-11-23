package client

import (
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

var (
	client *ClientSet
)

// SetGlobal todo
func SetGlobal(cli *ClientSet) {
	client = cli
}

// C Global
func C() *ClientSet {
	return client
}

// Client 客户端
type ClientSet struct {
	conn *grpc.ClientConn
	log  *zerolog.Logger
}

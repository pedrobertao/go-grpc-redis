package server

import (
	"fmt"
	"net"

	pb "github.com/pedrobertao/go-grpc-redis/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.CacheServiceServer
}

func Start() error {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	pb.RegisterCacheServiceServer(s, &server{})
	fmt.Println("gRPC is Listening")
	return s.Serve(listener)
}

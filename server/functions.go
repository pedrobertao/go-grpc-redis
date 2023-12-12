package server

import (
	"context"
	"net/http"

	pb "github.com/pedrobertao/go-grpc-redis/proto"
	redis "github.com/pedrobertao/go-grpc-redis/redis"
)

func (s *server) Health(ctx context.Context, in *pb.Empty) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{Message: http.StatusOK}, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	value, err := redis.Get(in.Key)
	if err != nil {
		return nil, err
	}
	return &pb.GetResponse{ResultCode: http.StatusOK, Result: value, ResultMessage: pb.HttpMessage_OK}, nil
}

func (s *server) Set(ctx context.Context, in *pb.SetRequest) (*pb.SetResponse, error) {
	err := redis.Set(in.Key, in.Value)
	if err != nil {
		return nil, err
	}
	return &pb.SetResponse{ResultCode: http.StatusOK, ResultMessage: pb.HttpMessage_OK}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := redis.Delete(in.Key)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{ResultCode: http.StatusOK, ResultMessage: pb.HttpMessage_OK, Result: "Delete " + in.Key}, nil
}

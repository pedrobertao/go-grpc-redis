package main

import (
	"log"

	redis "github.com/pedrobertao/go-grpc-redis/redis"
	grpc "github.com/pedrobertao/go-grpc-redis/server"
)

func main() {
	if err := redis.Start(); err != nil {
		log.Fatal(err)
	}

	if err := grpc.Start(); err != nil {
		log.Fatal(err)
	}
}

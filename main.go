package main

import (
	"context"
	"flag"
	"log"

	"github.com/go-redis/redis/v8"
)

func main() {
	serverName := flag.String("server", "127.0.0.1", "redis server (127.0.0.1 by default)")
	portNumber := flag.String("port", "6379", "redis port number (6379 by default)")
	databaseNumber := flag.Int("db", 0, "database number (0 by default)")

	flag.Parse()

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr: *serverName + ":" + *portNumber,
		DB:   *databaseNumber,
	})

	log.Println("Senging PING...")
	result, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	log.Println("Get:", result)
}

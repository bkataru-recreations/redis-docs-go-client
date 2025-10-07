package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func main() {
	// --- open a connection to redis ---

	opt, err := redis.ParseURL("redis://localhost:6379")
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	// Alternatively, one can also connect as
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // No password set
	// 	DB:       0,  // Use default DB
	// 	Protocol: 2,  // Connection protocol
	// })

	// --- test the connection by storing and retrieving a simple string ---

	ctx := context.Background()

	err = client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println("foo", val)

}

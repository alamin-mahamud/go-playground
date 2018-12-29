package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	ExampleNewCient()
}

func ExampleNewCient() {
	options := redis.Options{
		Addr: "localhost:6379",
		Password: "", // No Password Set
		DB: 0, // use default database
	}
	client := redis.NewClient(&options)
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
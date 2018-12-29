package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func main() {
	client, err := ExampleNewCient()
	if err != nil {
		panic(err)
	}

	// setting value
	err = settingValue(client)
	if err != nil {
		panic(err)
	}

	// getting value
	val, err := gettingValue(client)
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}

func ExampleNewCient() (*redis.Client, error){
	options := redis.Options{
		Addr: "localhost:6379",
		Password: "", // No Password Set
		DB: 0, // use default database
	}
	client := redis.NewClient(&options)
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func settingValue(client *redis.Client) error {
	return client.Set("Key", "value", 0).Err()
}

func gettingValue(client *redis.Client) (interface{}, error) {
	return client.Get("Key").Result()
}
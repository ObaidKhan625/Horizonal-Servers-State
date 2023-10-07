package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})
}

func setHandler(w http.ResponseWriter, r *http.Request) {

	err := redisClient.Set(context.Background(), "myKey", "Hello, Redis!", 0).Err()
	if err != nil {
		http.Error(w, "Failed to set data in Redis", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data set in Redis")
}

func getHandler(w http.ResponseWriter, r *http.Request) {

	val, err := redisClient.Get(context.Background(), "myKey").Result()
	if err != nil {
		http.Error(w, "Failed to get data from Redis", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Data from Redis: %s\n", val)
}

func main() {
	http.HandleFunc("/set", setHandler)
	http.HandleFunc("/get", getHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}

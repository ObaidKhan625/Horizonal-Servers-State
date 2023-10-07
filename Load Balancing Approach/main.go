package main

import (
	"net/http"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		c := cache.New(5*time.Minute, 10*time.Minute)
		serverName := os.Getenv("SERVER_NAME")

		// key := "key"
		// value := "value"

		// c.Set(key, value, cache.DefaultExpiration)

		cachedValue, found := c.Get("key")

		if found {
			stringValue, ok := cachedValue.(string)
			if ok {
				w.Write([]byte("Cached value: " + stringValue + " found in " + serverName))
			}
		} else {
			w.Write([]byte("Key not found in the cache of " + serverName))
		}

	})

	http.ListenAndServe(":8080", nil)
}

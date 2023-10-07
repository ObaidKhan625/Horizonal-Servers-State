package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("visit_count")
		serverName := os.Getenv("SERVER_NAME")
		if err != nil {
			http.SetCookie(w, &http.Cookie{
				Name:  "visit_count",
				Value: "1",
			})
			w.Write([]byte("Hello(x1), from " + serverName + "\n"))
			return
		}

		count, err := strconv.Atoi(cookie.Value)
		if err != nil {
			http.Error(w, "Invalid cookie value", http.StatusInternalServerError)
			return
		}

		count++
		cookie.Value = strconv.Itoa(count)
		http.SetCookie(w, cookie)

		responseMessage := fmt.Sprintf("Hello(x%d), from %s\n", count, serverName)
		w.Write([]byte(responseMessage))
	})

	http.ListenAndServe(":8080", nil)
}

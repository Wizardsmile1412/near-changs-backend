package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Near Changs API starting...")

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
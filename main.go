package main

import (
	"log"
	"server-bot/internal/server"
)

func main() {
	newServer := server.NewServer()

	err := newServer.ListenAndServe()
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}

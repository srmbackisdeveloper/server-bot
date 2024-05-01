package server

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"server-bot/internal/database"
)

type Server struct {
	port int
	db   database.Service
}

type APIServerError struct {
	Error string `json:"error"`
}

type APISuccessMessage struct {
	Message string `json:"message"`
}

func NewServer() *http.Server {
	port, err := strconv.Atoi("8080")
	if err != nil {
		log.Fatal(err)
	}

	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	server := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%d", port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	log.Printf("server running on port: %v\n", port)
	return server
}

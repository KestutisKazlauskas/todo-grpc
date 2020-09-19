package main

import (
	"log"

	"github.com/KestutisKazlauskas/todo-grpc/config"
	"github.com/KestutisKazlauskas/todo-grpc/server"
)

var (
	configPath = "config/config.yml"
)

func main() {
	// Loading configs
	conf, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Cann't load configurations, %v", err)
	}
	log.Printf("Configurations loaded %v", conf)

	//Starting a dabtabase here
	//And pass it to the models

	//Starting server
	srv, err := server.Create(conf)
	if err != nil {
		log.Fatalf("Could not create server from configs with error: %v", err)
	}
	log.Printf("Server %v", srv)
}

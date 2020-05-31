package main

import (
	"log"

	"github.com/KestutisKazlauskas/todo-grpc/config"
)

var (
	configPath = "config/config.yml"
)

func main() {
	conf, err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Cann't load configurations, %v", err)
	}
	log.Printf("Configurations loaded %v", conf)
}

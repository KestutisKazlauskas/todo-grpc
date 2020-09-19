package server

import (
	"fmt"

	"github.com/KestutisKazlauskas/todo-grpc/config"
)

// Handler is interface for server handeling
type Handler interface {
	Run() error
}

// Create is a function for creating server handler
func Create(conf *config.Config) (Handler, error) {

	switch conf.Server.Type {
	default:
		return nil, fmt.Errorf("Server type %v is not supported", conf.Server.Type)
	}

}

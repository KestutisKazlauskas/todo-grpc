package grpc

import (
	"github.com/KestutisKazlauskas/todo-grpc/config"
)

//Handler is a struct of grpc server
type Handler struct {
	Config *config.Config
}

// NewHandler is a function for creating new grpc handler for services
func NewHandler(conf *config.Config) (*Handler, error) {
	//TODO add the database int the domain folder and than us the domain logic here
	// Or could we use database layer here and domains left clean??
	handler := &Handler{
		Config: conf,
	}

	return handler, nil
}

// Run is a function to start grpc server with all services register
func (srv *Handler) Run() error {
	return nil
}

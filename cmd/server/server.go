package server

import (
	"net/http"
	"time"

	"delivery-chi-project/cmd/api"
	"delivery-chi-project/cmd/api/handler"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string, cartHandler *handler.CartHandler, inventoryHandler *handler.InventoryHandler) *Server {
	router := api.NewRouter(cartHandler, inventoryHandler)

	return &Server{
		httpServer: &http.Server{
			Addr:         ":" + port,
			Handler:      router,
			ReadTimeout:  15 * time.Second,
			WriteTimeout: 15 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Start() error {
	return s.httpServer.ListenAndServe()
}
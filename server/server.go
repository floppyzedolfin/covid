package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
)

type Server struct {
	app *fiber.App
	// data mapped on departement name
	data map[string][]entry
}

func New() *Server {
	s := Server{}
	s.data = make(map[string][]entry, 0)
	s.app = fiber.New()
	s.setupRoutes()
	return &s
}

func (s *Server) Listen(port int) {
	log.Fatal(s.app.Listen(fmt.Sprintf(":%d", port)))
}

func (s *Server) setupRoutes() {
	s.adminRoutes()
	s.apiRoutes()
	s.analyticsRoutes()
}

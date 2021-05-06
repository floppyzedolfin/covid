package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber"
)

// Server plugs onto the endpoints and can Listen() to a port
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

// setupRoutes sets up the routes (!) of the available endpoints
func (s *Server) setupRoutes() {
	// we have 3 kinds of routes, admin, api and analytics
	s.adminRoutes()
	s.apiRoutes()
	s.analyticsRoutes()
}

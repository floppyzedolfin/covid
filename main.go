package main

import (
	"github.com/floppyzedolfin/covid/server"
)

func main() {

	s := server.New()

	s.Listen(8405)
}

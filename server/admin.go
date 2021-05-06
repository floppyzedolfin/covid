package server

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber"
)

// adminRoutes adds the admin routes
func (s *Server) adminRoutes() {
	// GET to load data from source
	s.app.Get("/admin/load", s.load)
}

// gouvFile is a "global var" because I want to overwrite it in UTs
var gouvFile = "https://www.data.gouv.fr/fr/datasets/r/406c6a23-e283-4300-9484-54e78c8ae675"

// loads downloads the data from the remote server (gouv.fr) and stores it in this service
func (s *Server) load(c *fiber.Ctx) {
	// first download data
	resp, err := http.Get(gouvFile)
	if err != nil {
		panic(fmt.Sprintf("unable to download file %s: %s", gouvFile, err.Error()))
	}

	// parse the data
	r := csv.NewReader(resp.Body)
	r.Comma = ';'
	headersOn := true
	for {
		line, err := r.Read()
		if err == io.EOF {
			// end of file
			break
		}
		if err != nil {
			panic(fmt.Sprintf("unable to read entry: %s", err.Error()))
		}
		if headersOn {
			headersOn = false
			continue
		}
		if len(line) != 6 {
			panic(fmt.Sprintf("unvalid number of fields, expected 5, got %s", line))
		}
		// parse the entry
		var e entry
		e.State = line[0]
		e.Date = line[1]
		e.PositiveTests, err = strconv.Atoi(line[2])
		if err != nil {
			panic(fmt.Sprintf("invalid line %s, can't get positive tests: %s", line, err.Error()))
		}
		e.TestsPerformed, err = strconv.Atoi(line[3])
		if err != nil {
			panic(fmt.Sprintf("invalid line %s, can't get number of tests: %s", line, err.Error()))
		}
		e.AgeClass, err = strconv.Atoi(line[4])
		if err != nil {
			panic(fmt.Sprintf("invalid line %s, can't get age class: %s", line, err.Error()))
		}
		if len(s.data[e.State]) == 0 {
			s.data[e.State] = make([]entry, 1)
			s.data[e.State][0] = e
		} else {
			s.data[e.State] = append(s.data[e.State], e)
		}
	}
}

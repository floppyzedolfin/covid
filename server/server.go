package server

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber"
)



type Server struct {
	app *fiber.App
	// entries mapped on departement name
	entries map[string][]entry
}

type entry struct {
	Date string `json:"date"`
	State string `json:"departement"`
	TestsPerformed int `json:"testsPerformed"`
	AgeClass int `json:"ageClass"`
	PositiveTests int `json:"positiveTests"`

}

func New() *Server {
	s := Server{}
	s.entries = make(map[string][]entry, 0)
	s.app = fiber.New()
	err := s.load()
	if err != nil {
		panic(err)
	}
	s.setupRoutes()
	return &s
}

func (s *Server) Listen(port int) {
	s.app.Listen(fmt.Sprintf(":%d", port))
}

func (s *Server) setupRoutes() {
	// GET on all deps
	s.app.Get("/departements", s.DepartementsWithData)

	// GET on departement and data
	s.app.Get("/departement/:id/:date", s.DepartementWithDate)

	// POST on departements with range of dates
	s.app.Post("/departements", s.DepartementsWithDates)
}

func (s Server) DepartementsWithData(c *fiber.Ctx)  {
	departementIDs := make([]string, 0)
	for id := range s.entries {
		departementIDs = append(departementIDs, id)
	}

	c.JSON(departementIDs)
}

func (s Server) DepartementWithDate(c *fiber.Ctx) {
}

func (s Server) DepartementsWithDates(c *fiber.Ctx)  {
}


func (s *Server) load() error {
	// first download data
	fileURL := "https://www.data.gouv.fr/fr/datasets/r/406c6a23-e283-4300-9484-54e78c8ae675"
	resp, err := http.Get(fileURL)
	if err != nil {
		return fmt.Errorf("unable to download file %s: %w", fileURL, err)
	}

	r := csv.NewReader(resp.Body)
	headersOn := true
	for {
		line, err := r.Read()
		if err == io.EOF {
			// end of file
			break
		}
		if err != nil {
			return fmt.Errorf("unable to read entry: %w", err)
		}
		if headersOn {
			headersOn = false
			continue
		}
		if len(line) != 5 {
			return fmt.Errorf("unvalid number of fields, expected 5, got %d", len(line))
		}
		// parse the entry
		var e entry
		e.State = line[0]
		e.Date = line[1]
		e.TestsPerformed, err = strconv.Atoi(line[2])
		if err != nil {
			return fmt.Errorf("invalid line %s, can't get number of tests: %w", line, err)
		}
		e.AgeClass, err = strconv.Atoi(line[3])
		if err != nil {
			return fmt.Errorf("invalid line %s, can't get age class: %w", line, err)
		}
		e.PositiveTests, err = strconv.Atoi(line[4])
		if err != nil {
			return fmt.Errorf("invalid line %s, can't get positive tests: %w", line, err)
		}
		if len(s.entries[e.State]) == 0 {
			s.entries[e.State] = make([]entry, 1)
			s.entries[e.State][0] = e
		} else {
			s.entries[e.State] = append(s.entries[e.State], e)
		}
	}
	return nil
}

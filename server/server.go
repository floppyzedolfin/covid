package server

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber"
)

type Server struct {
	app *fiber.App
	// data mapped on departement name
	data map[string][]entry
}

type entry struct {
	Date           string `json:"date"`
	State          string `json:"departement"`
	TestsPerformed int    `json:"testsPerformed"`
	AgeClass       int    `json:"ageClass"`
	PositiveTests  int    `json:"positiveTests"`
}

func New() *Server {
	s := Server{}
	s.data = make(map[string][]entry, 0)
	s.app = fiber.New()
	err := s.load()
	if err != nil {
		panic(err)
	}
	s.setupRoutes()
	return &s
}

func (s *Server) Listen(port int) {
	log.Fatal(s.app.Listen(fmt.Sprintf(":%d", port)))
}

func (s *Server) setupRoutes() {
	// GET on all deps
	s.app.Get("/departements", s.DepartementsWithData)

	// GET on departement and data
	s.app.Get("/departement/:id/:date", s.DepartementWithDate)

	// POST on departements with range of dates
	s.app.Get("/departements/dates/:minDate/:maxDate", s.DepartementsWithDates)
}

func (s Server) DepartementsWithData(c *fiber.Ctx) {
	type departementsResponse struct {
		Departements []string `json:"departements"`
	}
	d := departementsResponse{Departements: make([]string, 0)}
	for id := range s.data {
		d.Departements = append(d.Departements, id)
	}

	c.JSON(d)
}

func (s Server) DepartementWithDate(c *fiber.Ctx) {
	type departementWithDateResponse struct {
		Data []entry `json:"data"`
	}
	dwd := departementWithDateResponse{Data: make([]entry, 0)}
	depData, ok := s.data[c.Params("id")]
	if !ok {
		// nothing found, return nothing
		c.JSON(dwd)
		return
	}
	date := c.Params("date")
	for _, e := range depData {
		if e.Date == date {
			dwd.Data = append(dwd.Data, e)
		}
	}
	c.JSON(dwd)
}

func (s Server) DepartementsWithDates(c *fiber.Ctx) {
	type departementsWithDateResponse struct {
		Data []entry `json:"data"`
	}
	dwd := departementsWithDateResponse{Data: make([]entry, 0)}
	minDate := c.Params("minDate")
	maxDate := c.Params("maxDate")
	for _, entries := range s.data {
		for _, entry := range entries {
			if dateBefore(minDate, entry.Date) && dateBefore(entry.Date, maxDate) {
				dwd.Data = append(dwd.Data)
			}
		}
	}
	return c.JSON(dwd)
}

func dateBefore(lhs, rhs string) bool {
	// we can compare strings :)
	layoutForm := "2006-01-02"
	tLeft, _ := time.Parse(layoutForm, lhs)
	tRight, _ := time.Parse(layoutForm, rhs)
	return tLeft.Before(tRight)
}

func (s *Server) load() error {
	// first download data
	fmt.Printf("downloading data")
	fileURL := "https://www.data.gouv.fr/fr/datasets/r/406c6a23-e283-4300-9484-54e78c8ae675"
	resp, err := http.Get(fileURL)
	if err != nil {
		return fmt.Errorf("unable to download file %s: %w", fileURL, err)
	}

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
			return fmt.Errorf("unable to read entry: %w", err)
		}
		if headersOn {
			headersOn = false
			continue
		}
		if len(line) != 6 {
			return fmt.Errorf("unvalid number of fields, expected 5, got %d", line)
		}
		// parse the entry
		var e entry
		e.State = line[0]
		e.Date = line[1]
		e.PositiveTests, err = strconv.Atoi(line[2])
		if err != nil {
			return fmt.Errorf("invalid line %s, can't get positive tests: %w", line, err)
		}
		e.TestsPerformed, err = strconv.Atoi(line[3])
		if err != nil {
			return fmt.Errorf("invalid line %s, can't get number of tests: %w", line, err)
		}
		e.AgeClass, err = strconv.Atoi(line[4])
		if err != nil {
			return fmt.Errorf("invalid line %s, can't get age class: %w", line, err)
		}
		if len(s.data[e.State]) == 0 {
			s.data[e.State] = make([]entry, 1)
			s.data[e.State][0] = e
		} else {
			s.data[e.State] = append(s.data[e.State], e)
		}
	}
	return nil
}

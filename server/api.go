package server

import (
	"sort"
	"strings"

	"github.com/gofiber/fiber"
)

// apiRoutes adds the API routes
func (s *Server) apiRoutes() {
	// GET on all deps
	s.app.Get("/api/departements", s.departementsWithData)

	// GET on departement and date
	s.app.Get("/api/departements/:id/dates/:date", s.departementDataWithDate)

	// GET on departements with range of dates
	s.app.Get("/api/departements/:ids/dates/:minDate/:maxDate", s.departementsWithDates)

	// GET on range of dates
	s.app.Get("/api/rangeDates", s.rangeDates)

	// GET on class ages
	s.app.Get("/api/classAges", s.classAges)
}

// departementsWithData returns the list of departements that contain data
func (s Server) departementsWithData(c *fiber.Ctx) {
	type departementsResponse struct {
		Departements []string `json:"departements"`
	}
	d := departementsResponse{Departements: make([]string, 0)}
	for id := range s.data {
		d.Departements = append(d.Departements, id)
	}
	sort.Strings(d.Departements)
	c.JSON(d)
}

// departementDataWithDate returns the data we have for a departement
func (s Server) departementDataWithDate(c *fiber.Ctx) {
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

// departementsWithDates returns the data we have for several
func (s Server) departementsWithDates(c *fiber.Ctx) {
	type departementsWithDateResponse struct {
		Data []entry `json:"data"`
	}
	dwd := departementsWithDateResponse{Data: make([]entry, 0)}
	departements := c.Params("ids")
	minDate := c.Params("minDate")
	maxDate := c.Params("maxDate")
	deps := strings.Split(departements, ",")
	for _, d := range deps {
		for _, entry := range s.data[d] {
			if dateBefore(minDate, entry.Date) && dateBefore(entry.Date, maxDate) {
				dwd.Data = append(dwd.Data, entry)
			}
		}
	}
	c.JSON(dwd)
}

func (s Server) classAges(c *fiber.Ctx) {
	type classAges struct {
		Ages []int `json:"ages"`
	}
	seenAges := make(map[int]struct{})
	for _, entries := range s.data {
		for _, e := range entries {
			if e.TestsPerformed > 0 {
				seenAges[e.AgeClass] = struct{}{}
			}
		}
	}

	agesList := make([]int, 0)
	for age := range seenAges {
		agesList = append(agesList, age)
	}
	sort.Ints(agesList)
	c.JSON(classAges{Ages: agesList})
}

func (s Server) rangeDates(c *fiber.Ctx) {
	type rangeDatesResponse struct {
		MinDate string `json:"minDate"`
		MaxDate string `json:"maxDate"`
	}
	var r rangeDatesResponse
	for _, entries := range s.data {
		for _, entry := range entries {
			if r.MinDate == "" {
				r.MinDate = entry.Date
				r.MaxDate = entry.Date
				continue
			}
			switch {
			case dateBefore(entry.Date, r.MinDate):
				r.MinDate = entry.Date
			case dateBefore(r.MaxDate, entry.Date):
				r.MaxDate = entry.Date
			}
		}
	}
	c.JSON(r)
}

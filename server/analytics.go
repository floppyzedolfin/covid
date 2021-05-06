package server

import "github.com/gofiber/fiber"

// analyticsRoutes adds the analytics routes
func (s *Server) analyticsRoutes() {
	// GET national data
	s.app.Get("/analytics/national/dates/:minDate/:maxDate", s.nationalWithDates)

	// GET departement brief info
	s.app.Get("/analytics/departements/:id/brief", s.departementBrief)

	// GET topDepartements
	s.app.Get("/analytics/dates/:date/top5", s.top5)
}

// nationalWithDates returns statistics at national level
func (s Server) nationalWithDates(c *fiber.Ctx) {
	type nationalData struct {
		TestsNumbers  int     `json:"testsNumber"`
		PositiveTests int     `json:"positiveTests"`
		Ratio         float32 `json:"ratio"`
	}

	n := nationalData{}
	minDate := c.Params("minDate")
	maxDate := c.Params("maxDate")
	for _, entries := range s.data {
		for _, entry := range entries {
			if dateBefore(minDate, entry.Date) && dateBefore(entry.Date, maxDate) {
				n.TestsNumbers += entry.TestsPerformed
				n.PositiveTests += entry.PositiveTests
			}
		}
	}
	n.Ratio = float32(n.PositiveTests) / float32(n.TestsNumbers)
	c.JSON(n)
}

// departementBrief returns a short version of the details for a departement over time
func (s Server) departementBrief(c *fiber.Ctx) {
	type departementBriefResponse struct {
		DateMostTests    string
		DateMostPositive string
		DateHightestRate string
	}
	dr := departementBriefResponse{}
	depData, ok := s.data[c.Params("id")]
	if !ok {
		// nothing found, return nothing
		c.JSON(dr)
		return
	}
	highestTests := 0
	highestPositives := 0
	highestRate := 0.0
	currDate := depData[0].Date
	currDateTests := 0
	currDatePositives := 0
	for _, e := range depData {
		if e.Date != currDate {
			// it's a new date - check if we have to update things
			if currDateTests > 10 {
				if currDateTests > highestTests {
					highestTests = currDateTests
					dr.DateMostTests = currDate
				}
				if currDatePositives > highestPositives {
					highestPositives = currDatePositives
					dr.DateMostPositive = currDate
				}
				currDateRate := float64(currDatePositives) / float64(currDateTests)
				if currDateRate > highestRate {
					highestRate = currDateRate
					dr.DateHightestRate = currDate
				}
			}
			currDate = e.Date
			currDateTests = 0
			currDatePositives = 0
		}
		currDateTests += e.TestsPerformed
		currDatePositives += e.PositiveTests
	}

	// don't forget last day
	if currDateTests > 10 {
		if currDateTests > highestTests {
			highestTests = currDateTests
			dr.DateMostTests = currDate
		}
		if currDatePositives > highestPositives {
			highestPositives = currDatePositives
			dr.DateMostPositive = currDate
		}
		currDateRate := float64(currDatePositives) / float64(currDateTests)
		if currDateRate > highestRate {
			highestRate = currDateRate
			dr.DateHightestRate = currDate
		}
	}

	c.JSON(dr)
}


// top5 returns the top5 departements for specific ratios
func (s Server) top5(c *fiber.Ctx) {
	// TODO
	c.JSON(struct{}{})
}
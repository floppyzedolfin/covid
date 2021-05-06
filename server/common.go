package server

import (
	"time"
)

// entry is a line in the CSV (without the population information)
type entry struct {
	Date           string `json:"date"`
	State          string `json:"departement"`
	TestsPerformed int    `json:"testsPerformed"`
	AgeClass       int    `json:"ageClass"`
	PositiveTests  int    `json:"positiveTests"`
}

// dateBefore compares two dates
func dateBefore(lhs, rhs string) bool {
	// we can compare strings :)
	layoutForm := "2006-01-02"
	tLeft, _ := time.Parse(layoutForm, lhs)
	tRight, _ := time.Parse(layoutForm, rhs)
	return tLeft.Before(tRight)
}

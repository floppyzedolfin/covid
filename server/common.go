package server

import (
	"time"
)

type entry struct {
	Date           string `json:"date"`
	State          string `json:"departement"`
	TestsPerformed int    `json:"testsPerformed"`
	AgeClass       int    `json:"ageClass"`
	PositiveTests  int    `json:"positiveTests"`
}

func dateBefore(lhs, rhs string) bool {
	// we can compare strings :)
	layoutForm := "2006-01-02"
	tLeft, _ := time.Parse(layoutForm, lhs)
	tRight, _ := time.Parse(layoutForm, rhs)
	return tLeft.Before(tRight)
}

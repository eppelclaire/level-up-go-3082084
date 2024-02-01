package main

import (
	"flag"
	"log"
	"time"
)

const expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) time.Time {
	formattedTarget, err := time.Parse(expectedFormat, target)
	if err != nil || time.Now().After(formattedTarget) {
		panic(err)
	}
	return formattedTarget
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
	return time.Until(target).Hours() / 24
}

func main() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()
	target := parseTime(*bday)
	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}

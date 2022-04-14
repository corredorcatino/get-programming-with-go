// Ticket to Mars
package main

import (
	"fmt"
	"math/rand"
)

const distanceToMars = 62100000 // km
const secondsPerDay = 86400

func main() {
	fmt.Printf("%-16v %-4v %-10v %-4v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Printf("======================================\n")

	for count := 0; count < 10; count++ {
		spaceline := "Space Adventures"
		switch sp := rand.Intn(3); sp {
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Galactic"
		}

		speed := rand.Intn(15) + 16 // km/s
		duration := (distanceToMars / speed) / secondsPerDay

		tripType := "One-Way"
		switch tt := rand.Intn(2); tt {
		case 1:
			tripType = "Round-trip"
		}

		price := speed + 20 // In millions of dollars
		if tripType == "Round-trip" {
			price *= 2
		}

		fmt.Printf("%-16v %4v %-10v $%4v\n", spaceline, duration, tripType, price)
	}
}

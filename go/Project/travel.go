package main

import (
	"fmt"
)

func fuelLeft(fuel int) {
	fmt.Println("Still you have", fuel, "gallons fuel left !!")
}

func cantFly() {
	fmt.Println("Sorry, You don't have enough fuel to go ..")
}

func planetGreet(planet string) {
	fmt.Println("Welcome to", planet, "!!")
}

func calculateFuel(planet string) int {

	var fuel int

	switch planet {
	case "Mercury":
		fuel = 100000
	case "Venus":
		fuel = 200000
	case "Earth":
		fuel = 300000
	case "Mars":
		fuel = 400000
	case "Jupiter":
		fuel = 500000
	case "Saturn":
		fuel = 600000
	case "Uranus":
		fuel = 700000
	case "Neptune":
		fuel = 800000
	case "Pluto":
		fuel = 900000
	default:
		fmt.Println("Enter a valid Planet!!")
	}
	return fuel
}

func flyToPlanet(planet string, fuel int) int {
	remainingFuel := fuel
	fuelCost := calculateFuel(planet)

	if remainingFuel >= fuelCost {
		planetGreet(planet)
		remainingFuel -= fuelCost
	} else {
		cantFly()
	}
	return fuel

}
func main() {
	fuel := 1000000
	planet := "Mercury"

	fuel = flyToPlanet(planet, fuel)
	fuelLeft(fuel)

}

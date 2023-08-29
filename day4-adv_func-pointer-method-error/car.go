package main

import "fmt"

type Car struct {
	Type    string
	FuelInL float64
}

func (c Car) EstimateDistance() float64 {
	fuelEfficiency := 1.5 // L/mill
	estimatedDistance := c.FuelInL * fuelEfficiency
	return estimatedDistance
}

func main() {
	car := Car{
		Type:    "Sedan",
		FuelInL: 30.0,
	}

	estimatedDistance := car.EstimateDistance()
	fmt.Printf("Estimated distance for the %s car: %.2f kilometers\n", car.Type, estimatedDistance)
}

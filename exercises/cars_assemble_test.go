package exercises

import "testing"

func TestCalculateWorkingCarsPerHour(t *testing.T) {
	carRate := 1200
	carEfficiency := 95.0
	hourlyProductionRate := CalculateWorkingCarsPerHour(carRate, carEfficiency)

	if hourlyProductionRate != 1140.0 {
		t.Error()
	}
}

func TestCalculateWorkingCarsPerMinute(t *testing.T) {
	carRate := 1200
	carEfficiency := 95.0
	minuteProduction := CalculateWorkingCarsPerMinute(carRate, carEfficiency)

	if minuteProduction != 19 {
		t.Error()
	}
}

func TestCarAssemblyMultipleBatches(t *testing.T) {
	numberOfCars := 37
	totalCost := CalculateCost(numberOfCars)

	if totalCost != 355000 {
		t.Error()
	}
}

func TestCarAssemblyLessThanABatch(t *testing.T) {
	numberOfCars := 9
	totalCost := CalculateCost(numberOfCars)

	if totalCost != 90000 {
		t.Error()
	}
}

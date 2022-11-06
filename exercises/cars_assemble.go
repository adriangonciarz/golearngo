package exercises

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100.0
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

func CalculateCost(carsCount int) uint {
	batches := int(carsCount / 10)
	remaining := carsCount % 10

	return uint(batches*95000 + remaining*10000)
}

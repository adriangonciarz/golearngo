// https://exercism.org/tracks/go/exercises/lasagna
package exercises

const OvenTime int = 40

func RemainingOvenTime(actual int) int {
	return OvenTime - actual
}

func PreparationTime(numberOfLayers int) int {
	// Each layer takes 2 minutes to prepare
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}

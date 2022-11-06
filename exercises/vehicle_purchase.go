package exercises

import (
	"fmt"
	"strings"
)

func NeedsLicense(kind string) bool {
	switch kind {
	case "car", "truck":
		return true
	}
	return false
}

func ChooseVehicle(option1, option2 string) string {
	var betterOption string
	if strings.Compare(option1, option2) == -1 {
		betterOption = option1
	} else {
		betterOption = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", betterOption)
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3.0 {
		return originalPrice * 0.8
	} else if age < 10.0 {
		return originalPrice * 0.7
	} else {
		return originalPrice * 0.5
	}
}

package exercises

import "testing"

func TestSumTwoIntegers(t *testing.T) {
	total := sumTwoIntegers(3, 5)
	if total != 8 {
		t.Error()
	}
}

func TestZeroCelsiusToFahrenheit(t *testing.T) {
	zeroInFahrenheit := celsiusToFahrenheit(0)
	if zeroInFahrenheit != 32 {
		t.Error()
	}
}

func TestNegativeCelsiusToFahrenheit(t *testing.T) {
	zeroInFahrenheit := celsiusToFahrenheit(-20.0)
	if zeroInFahrenheit != 0 {
		t.Error()
	}
}

func PositiveTestCelsiusToFahrenheit(t *testing.T) {
	zeroInFahrenheit := celsiusToFahrenheit(20.0)
	if zeroInFahrenheit != 64 {
		t.Error()
	}
}

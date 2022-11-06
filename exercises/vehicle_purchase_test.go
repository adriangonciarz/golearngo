package exercises

import "testing"

func TestNeedsLicense(t *testing.T) {
	tables := []struct {
		kind         string
		needsLicense bool
	}{
		{"car", true},
		{"truck", true},
		{"bike", false},
		{"moto", false},
	}

	for _, table := range tables {
		if NeedsLicense(table.kind) != table.needsLicense {
			t.Error()
		}
	}
}

func TestChooseVehicle(t *testing.T) {
	tables := []struct {
		optionA string
		optionB string
		result  string
	}{
		{"Ford", "Toyota", "Ford is clearly the better choice."},
		{"Ford", "Bentley", "Bentley is clearly the better choice."},
		{"Audi", "Audi", "Audi is clearly the better choice."},
	}

	for _, table := range tables {
		if ChooseVehicle(table.optionA, table.optionB) != table.result {
			t.Error()
		}
	}

}

func TestCalculateVehiclePrice(t *testing.T) {
	tables := []struct {
		originalPrice float64
		age           float64
		currentPrice  float64
	}{
		{10500.0, 2.0, 8400.0},
		{10500.0, 3.0, 7350.0},
		{10500.0, 10.0, 5250.0},
	}

	for _, table := range tables {
		if CalculateResellPrice(table.originalPrice, table.age) != table.currentPrice {
			t.Error()
		}
	}

}

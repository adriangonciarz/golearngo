package exercises

import "testing"

func TestNonErrorneousHummingDistance(t *testing.T) {

	tables := []struct {
		stringA  string
		stringB  string
		distance int
	}{
		{"CGACGTATGCA", "CGACGGATGGA", 2},
		{"", "", 0},
		{"CGACGTATGCA", "CGACGTATGCA", 0},
		{"BBBBBBBBBBB", "CGACGTATGCA", 11},
	}

	for _, table := range tables {
		actualDistance, _ := Distance(table.stringA, table.stringB)
		if actualDistance != table.distance {
			t.Error()
		}
	}
}

func TestErroneousHummingDistance(t *testing.T) {
	stringA := "ABC"
	stringB := "ABCA"

	_, actualError := Distance(stringA, stringB)
	if actualError == nil {
		t.Error()
	}
}

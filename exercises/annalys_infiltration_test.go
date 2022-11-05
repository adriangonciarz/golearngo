package exercises

import "testing"

func TestCanFastAttack(t *testing.T) {
	tables := []struct {
		knightIsAwake bool
		canFastAttack bool
	}{
		{true, false},
		{false, true},
	}

	for _, table := range tables {
		canFastAttack := CanFastAttack(table.knightIsAwake)
		if canFastAttack != table.canFastAttack {
			t.Error()
		}
	}
}

func TestCanSpy(t *testing.T) {
	tables := []struct {
		knightIsAwake   bool
		archerIsAwake   bool
		prisonerIsAwake bool
		canSpy          bool
	}{
		{true, false, false, true},
		{false, true, false, true},
		{false, false, true, true},
		{false, false, false, false},
	}

	for _, table := range tables {
		canSpy := CanSpy(table.knightIsAwake, table.archerIsAwake, table.prisonerIsAwake)
		if canSpy != table.canSpy {
			t.Error()
		}
	}
}

func TestCanSignalPrisoner(t *testing.T) {
	tables := []struct {
		archerIsAwake   bool
		prisonerIsAwake bool
		canSignal       bool
	}{
		{true, false, false},
		{false, true, true},
		{false, false, false},
		{true, true, false},
	}

	for _, table := range tables {
		canSpy := CanSignalPrisoner(table.archerIsAwake, table.prisonerIsAwake)
		if canSpy != table.canSignal {
			t.Error()
		}
	}
}

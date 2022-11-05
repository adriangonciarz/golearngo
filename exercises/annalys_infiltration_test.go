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

func TestCanFreePrisoner(t *testing.T) {
	tables := []struct {
		knightIsAwake   bool
		archerIsAwake   bool
		prisonerIsAwake bool
		hasHerDog       bool
		canFreePrisoner bool
	}{
		{true, false, true, true, true},
		{true, false, false, true, true},
		{false, false, true, false, true},
		{false, false, true, true, true},

		{false, false, false, false, false},
		{true, true, true, false, false},
	}

	for _, table := range tables {
		canSpy := CanFreePrisoner(table.knightIsAwake, table.archerIsAwake, table.prisonerIsAwake, table.hasHerDog)
		if canSpy != table.canFreePrisoner {
			t.Error()
		}
	}
}

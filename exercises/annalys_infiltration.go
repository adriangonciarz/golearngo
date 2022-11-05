// https://exercism.org/tracks/go/exercises/annalyns-infiltration
package exercises

func CanFastAttack(knightIsAwake bool) bool {
	if knightIsAwake {
		return false
	} else {
		return true
	}
}

func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	if knightIsAwake || archerIsAwake || prisonerIsAwake {
		return true
	} else {
		return false
	}
}

func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	if !archerIsAwake && prisonerIsAwake {
		return true
	} else {
		return false
	}
}

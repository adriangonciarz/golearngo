package exercises

func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

func FirstTurn(card1, card2, dealerCard string) string {
	mySum := (ParseCard(card1) + ParseCard(card2))
	dealerValue := ParseCard(dealerCard)
	if mySum <= 11 {
		return "H"
	} else if 12 <= mySum && mySum <= 16 && dealerValue >= 7 {
		return "H"
	} else if mySum == 21 && dealerValue < 10 {
		return "W"
	} else if card1 == "ace" && card2 == "ace" {
		return "P"
	} else {
		return "S"
	}
}

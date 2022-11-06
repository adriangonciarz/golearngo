package exercises

import (
	"fmt"
	"strings"
)

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*", numStarsPerLine)
	return fmt.Sprintf("%s\n%s\n%s", border, welcomeMsg, border)
}

func CleanupMessage(oldMsg string) string {
	msg := strings.ReplaceAll(oldMsg, "*", "")
	msg = strings.ReplaceAll(msg, "\n", "")
	return strings.Trim(msg, " ")
}

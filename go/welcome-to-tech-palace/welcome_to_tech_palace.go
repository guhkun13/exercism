package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer) 
	// panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// panic("Please implement the AddBorder() function")

	var stars string;
	for i := 0; i < numStarsPerLine; i++ {
		stars += "*"
	}
	msg := fmt.Sprintf("%s\n%s\n%s", stars, welcomeMsg, stars)
	return msg
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// panic("Please implement the CleanupMessage() function")
	newMsg := strings.ReplaceAll(oldMsg, "*", "")
	fmt.Println(newMsg)

	newMsg = strings.TrimSpace(newMsg)
	
	return newMsg
}

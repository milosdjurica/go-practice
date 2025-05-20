package greetings

import "fmt"

func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	// Upper code is shortcut for 2 lines below
	// var message string
	// message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

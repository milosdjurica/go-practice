package main

import (
	"fmt"

	greetings "github.com/milosdjurica/go-practice/02-greetings"
)

func main() {
	message := greetings.Hello("Milos")
	fmt.Println(message)
}

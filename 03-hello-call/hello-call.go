package main

import (
	"fmt"
	"log"

	greetings "github.com/milosdjurica/go-practice/02-greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Milos")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	names := []string{"Milos", "Decenter", "Solidity", "Golang"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

}

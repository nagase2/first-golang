package main

import (
	"fmt"
	"log"
	"example.com/greetings"
)

func main() {
	// Get a greeting message and print it.
	message, err := greetings.Hello("sss")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}

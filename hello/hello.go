package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// ここで設定することにより先頭のタイムスタンプがなくなり、代わりに指定文字列が入る
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// Get a greeting message and print it.
	//message, err := greetings.Hello("aaa")
	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)


	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

	m := make(map[string]int)
	m["aaa"] = 23
	m["sss"] = 232
	m["aaa"] = 231
	fmt.Println(m["ii"])
}

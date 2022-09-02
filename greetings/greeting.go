package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// If no name was given, throw error
	if name == "" {
		return "", errors.New("emppty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil

}

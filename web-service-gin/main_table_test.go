package main

import (
	"fmt"
	"testing"
)

func TestSomething(t *testing.T) {
	testTable := []struct {
		name string
		val1 int
	}{
		{
			name: "aaaa",
			val1: 23,
		},
		{
			name: "bbb",
			val1: 23,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name, tt.val1)

		})
	}
}

func TestSomething2(t *testing.T) {
	testTable := []struct {
		name string
		val1 int
	}{
		{
			name: "ああああ",
			val1: 23,
		},
		{
			name: "bbb",
			val1: 23,
		},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(tt.name, tt.val1)

		})
	}
}

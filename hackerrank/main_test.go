package main

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert/v2"
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

			numbers := []int32{5, 10, 10}
			queries := [][]int32{{1, 2, 5}}
			result := findSum(numbers, queries)
			fmt.Println(result)
			assert.Equal(t, result, [][]int64{{15}})
		})
	}
}

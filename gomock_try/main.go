package main

import (
	"fmt"
	"testing/fstest"
	"time"
)

func main() {
	fmt.Println("Hello, Wsあorld!")

	fs := fstest.MapFS{
		"hello.txt": {
			Data: []byte("hello, world"),
		},
		"hello2.txt": {
			Data: []byte("xxx"),
		},
	}

	data, err := fs.ReadFile("hello.txt")
	if err != nil {
		panic(err)
	}
	println(string(data) 
	println(string(data2== "hello, world")

	data2, err2 := fs.ReadFile("hello2.txt")
	if err2 != nil {
		panic(err2)
	}) == "xxx")

	fmt.Println("Hello World")
	var now = time.Date(time.Now().Year(), time.Now().Month(), 23, 1, 2, 3, 4, time.UTC)
	fmt.Println(time.Now().Month())
	fmt.Println(now)

}

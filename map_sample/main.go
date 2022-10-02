package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

type Car interface {
	getName() string
	getColor() string
}
type Vechle interface {
	run() string
}

type Bot interface {
	getGreeting() string
	printGreeting()
}

type EnglishBot struct {
	greeting string
}

func (b *EnglishBot) getGreeting() string {
	return b.greeting
}
func (b *EnglishBot) printGreeting() {
	b.greeting += "b"
	fmt.Println("Hello!!", b.greeting)
}

type SpanishBot struct {
	greeting string
}

func (b *SpanishBot) getGreeting() string {
	return b.greeting
}
func (b *SpanishBot) printGreeting() {
	//fmt.Println(b.greeting)
	b.greeting += "a"
	fmt.Println("Ora!!", b.greeting)
}

func printGreeting(b Bot) {
	b.printGreeting()
}

func printCarName(c Car) {
	fmt.Println(c.getName())
}

type ToyotaCar struct{}

// run implements Vechle
func (*ToyotaCar) run() string {
	message := "run toyota!"
	fmt.Println(message)
	return message
}
func (t *ToyotaCar) getName() string {
	return "toyota desu"
}
func (t *ToyotaCar) getColor() string {
	return "green"
}

type HondaCar struct{}

func (h HondaCar) getName() string {
	return "honda desu"
}
func (h HondaCar) getColor() string {
	return "red"
}

var flagvar int

func init() {

}
func main() {
	wordPtr := flag.String("word", "foo", "a string")
	boolPtr := flag.Bool("fork", false, "a bool")
	debugPtr := flag.Bool("debug", false, "debug flag. if  it's true")
	flag.Parse()
	fmt.Println("isDebug:", *debugPtr)

	//os.Stdout = nil    // turn it off
	fmt.Println("word:", *wordPtr)
	fmt.Println("fork:", *boolPtr)
}

func test1() {
	myMap := map[string]string{
		"aa": "11",
		"bb": "22",
		"cc": "33",
	}
	fmt.Println(myMap)
	for name, value := range myMap {
		fmt.Print("😁", name)
		fmt.Println(" 🏰", value)
	}

	var sb = SpanishBot{}
	eb := EnglishBot{}
	var ca Car = &ToyotaCar{}
	fmt.Println("🚗", ca.getName())

	ca = &HondaCar{}
	fmt.Println("🚙", ca.getName())

	var vc Vechle = &ToyotaCar{}
	vc.run()

	printGreeting(&sb)
	printGreeting(&eb)
	printGreeting(&eb)

	tcar := ToyotaCar{}
	//この行を書くことでToyotaCarがCarのインターフェースメンバーに追加されたことになる？
	printCarName(&tcar)

	// http
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// write pattern 1
	bs := make([]byte, 111)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}

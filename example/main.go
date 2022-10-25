package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct {}
type spanishBot struct {}


func (englishBot) getGreeting() string {
	return "Hi!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreating(b bot) {
	fmt.Println(b.getGreeting())

}

func main() {
	var sb spanishBot
	var eb englishBot
	
	printGreating(sb)
	printGreating(eb)
}

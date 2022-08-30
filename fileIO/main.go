package main

import "fmt"


func main() {
	cards := newDeck()
	cards.print()
	hand, cards := deal(cards, 2)
	fmt.Println(hand.toString())
	fmt.Println(cards.toString())
	cards.saveToFile("cards.csv")
	newCards := newDeckFromFile("cards.csv")
	newCards.print()
}
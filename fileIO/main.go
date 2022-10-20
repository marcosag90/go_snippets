package main

import "fmt"


func main() {
	cards := newDeck()
	cards.print()
	hand, cards := deal(cards, 2)
	fmt.Println("Hand")
	hand.print()
	fmt.Println("Remaining deck")
	cards.print()
	cards.saveToFile("cards.csv")
	newCards := newDeckFromFile("cards.csv")
	newCards.print()	
	fmt.Println("Shufling cards...")
	newCards.shuffle()
	newCards.print()
}
package main

import (
	"math/rand"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards :=deck{} // slice
	suits := []string{"Spades", "Hearts"}
	values := []string{"One", "Two", "Three"}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value + " of " + suit )
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}  

func newDeckFromFile(filename string) deck {
	d, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: " + err.Error() )
		os.Exit(1);
	}
	s := strings.Split(string(d), ", ")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) -1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
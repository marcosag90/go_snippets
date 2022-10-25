package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type card struct {
	suit string
	value string
}

func (c card) toString() string {
	return c.value + " of " + c.suit 
}

func (c card) print() {
	fmt.Println(c.toString());
}


type deck []card

func newDeck() deck {
	cards := deck{} // slice
	suits := []string{"Spades", "Hearts"}
	values := []string{"One", "Two", "Three"}

	for _, s := range suits {
		for _, v := range values {
			cards = append(cards, card{ value: v, suit: s } )
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card.toString())
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	var ds []string
	for _, c := range d {
		ds = append(ds, c.toString())
	}
	return strings.Join(ds, ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}  

func newDeckFromFile(filename string) deck {
	s, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: " + err.Error() )
		os.Exit(1);
	}
	ss := strings.Split(string(s), ", ")
	var d deck
	for _, cs := range ss {
		v := strings.Split( cs, " of ")
		d = append(d, card{ value: v[0], suit: v[1] })
	}
	return d
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	
	for i := range d {
		newPosition := r.Intn(len(d) -1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
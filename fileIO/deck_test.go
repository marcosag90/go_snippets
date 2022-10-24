package main

import (
	"errors"
	"os"
	"testing"
)

var s = 2 // number of suits
var v = 3 // number of values

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != s*v {
		t.Errorf("Expected deck lenght of %d, bug got %d", s*v, len(d))
	}

	if d[0] != "One of Spades" {
		t.Errorf("Expected first card to be One of Spades but got %s", d[0])
	}

	if d[len(d)-1] != "Three of Hearts" {
		t.Errorf("Expected last card to be Three of Hearts but got %s", d[len(d)-1])
	}
}

func TestSaveDeckToFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	if _, err := os.Stat("_decktesting"); errors.Is(err, os.ErrNotExist) {
		t.Errorf("Deck of cards could not create a new file fo store deck in path")
	}
	os.Remove("_decktesting")
}

func TestLoadDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	
	d = newDeckFromFile("_decktesting")
	
	if len(d) != s*v {
		t.Errorf("Expected deck lenght of %d, bug got %d", s*v, len(d))
	}

	if d[0] != "One of Spades" {
		t.Errorf("Expected first card to be One of Spades but got %s", d[0])
	}

	if d[len(d)-1] != "Three of Hearts" {
		t.Errorf("Expected last card to be Three of Hearts but got %s", d[len(d)-1])
	}
	os.Remove("_decktesting")
}
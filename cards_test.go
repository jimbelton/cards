package main

import "./cards"
import "testing"

var ranks = map[string] bool {
    "A":  true,
    "2":  true,
    "3":  true,
    "4":  true,
    "5":  true,
    "6":  true,
    "7":  true,
    "8":  true,
    "9":  true,
    "10": true,
    "J":  true,
    "Q":  true,
    "K":  true,
}
var suits = map[string] bool {
    "hearts":   true,
    "spades":   true,
    "clubs":    true,
    "diamonds": true,
}

func TestDeck(t *testing.T) {
    deck := cards.NewDeck()
    card := deck.DealCard()

    if !ranks[card.Rank] {
        t.Errorf("Unexpected rank %s", card.Rank)
    }

    if !suits[card.Suit] {
        t.Errorf("Unexpected suit %s", card.Suit)
    }
}

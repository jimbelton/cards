package main

import "./blackjack"
import "./cards"
import "testing"

func TestHand(t *testing.T) {
    deck := cards.NewDeck()
    _ = blackjack.NewHand(deck)
}
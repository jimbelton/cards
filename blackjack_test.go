package main

import "./blackjack"
import "./cards"
import "testing"

func TestHand(t *testing.T) {
    deck := cards.NewDeck()
    hand := blackjack.NewHand(deck)

    if hand.NumCards() != 2 {
        t.Errorf("A blackjack hand should have 2 cards, not %d", hand.NumCards())
    }
}
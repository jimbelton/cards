package main

import "./blackjack"
import "./cards"
import "testing"

func TestHand(t *testing.T) {
    deck := cards.NewDeck()
    hand := blackjack.NewHand(deck)

    deck.Stack(cards.NewCard("2", "hearts"))
    deck.Stack(cards.NewCard("2", "spades"))

    if hand.Score() != 4 {
        t.Errorf("Blackjack hand %s should have a score of 4, not %d", hand.ToStr(), hand.Score())
    }
}
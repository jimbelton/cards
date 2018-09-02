package main

import "./blackjack"
import "./cards"
import "testing"

func TestHand(t *testing.T) {
    deck := cards.NewDeck()

    deck.Stack(cards.NewCard("2", "hearts"))
    deck.Stack(cards.NewCard("2", "spades"))
    hand := blackjack.NewHand(deck)

    if hand.Score() != 4 {
        t.Errorf("Blackjack hand %s should have a score of 4, not %d", hand.ToStr(), hand.Score())
    }

    deck.Stack(cards.NewCard("10", "diamonds"))
    deck.Stack(cards.NewCard("10", "clubs"))
    hand = blackjack.NewHand(deck)

    if hand.Score() != 20 {
        t.Errorf("Blackjack hand %s should have a score of 20, not %d", hand.ToStr(), hand.Score())
    }
}
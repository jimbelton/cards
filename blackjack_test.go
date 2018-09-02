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

    deck.Stack(cards.NewCard("Q", "hearts"))
    deck.Stack(cards.NewCard("A", "spades"))
    hand = blackjack.NewHand(deck)

    if hand.Score() != 21 {
        t.Errorf("Blackjack hand %s should have a score of 21, not %d", hand.ToStr(), hand.Score())
    }

    if hand.CanSplit() {
        t.Errorf("Blackjack hand %s should not be splitable", hand.ToStr())
    }

    deck.Stack(cards.NewCard("A", "diamonds"))
    deck.Stack(cards.NewCard("A", "clubs"))
    hand = blackjack.NewHand(deck)

    if hand.Score() != 12 {
        t.Errorf("Blackjack hand %s should have a score of 12, not %d", hand.ToStr(), hand.Score())
    }

    if !hand.CanSplit() {
        t.Errorf("Blackjack hand %s should be splitable", hand.ToStr())
    }
}
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

    deck.Stack(cards.NewCard("K", "diamonds"))
    deck.Stack(cards.NewCard("9", "clubs"))

    split := hand.Split()

    if hand.Score() != 20 {
        t.Errorf("Blackjack hand %s should have a score of 20, not %d", hand.ToStr(), hand.Score())
    }

    if split.Score() != 21 {
        t.Errorf("Blackjack hand %s should have a score of 21, not %d", split.ToStr(), split.Score())
    }

    deck.Stack(cards.NewCard("A", "spades"))
    hand.Hit()

    if hand.Score() != 21 {
        t.Errorf("Blackjack hand %s (after hit) should have a score of 21, not %d", hand.ToStr(), hand.Score())
    }
}

func TestDeckExhaustionDealingHand(t *testing.T) {
    deck := cards.NewDeck()

    for i := 0; i < 51; i++ {
        deck.DealCard()
    }

    if (blackjack.NewHand(deck) != nil) {
        t.Errorf("Dealt a hand when deck had only 1 card left")
    }

    deck.Stack(cards.NewCard("A", "diamonds"))
    deck.Stack(cards.NewCard("A", "clubs"))
    hand := blackjack.NewHand(deck)

    if (hand == nil) {
        t.Errorf("Failed to deal a hand when deck had 2 cards")
    } else if hand.Hit() != nil {
         t.Errorf("Hit when deck had no cards left")
    }
}
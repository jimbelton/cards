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

    if hand.IsBlackjack() {
        t.Errorf("Blackjack hand %s should not be a blackjack", hand.ToStr())
    }

    deck.Stack(cards.NewCard("Q", "hearts"))
    deck.Stack(cards.NewCard("A", "spades"))
    hand = blackjack.NewHand(deck)

    if hand.ScoreCard(0) != 11 {
        t.Errorf("Blackjack ace %s should have a score of 11, not %d", hand.Cards[0].ToStr(), hand.ScoreCard(0))
    }

    if hand.ScoreCard(1) != 10 {
        t.Errorf("Blackjack queen %s should have a score of 10, not %d", hand.Cards[1].ToStr(), hand.ScoreCard(1))
    }

    if hand.Score() != 21 {
        t.Errorf("Blackjack hand %s should have a score of 21, not %d", hand.ToStr(), hand.Score())
    }

    if !hand.IsBlackjack() {
        t.Errorf("Blackjack hand %s should be a blackjack", hand.ToStr())
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

    if hand.IsBlackjack() {
        t.Errorf("Blackjack hand %s should not be a blackjack", hand.ToStr())
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

    deck.Stack(cards.NewCard("K", "clubs"))

    if !hand.CanSplit() {
        t.Errorf("Can't split a hand with two aces")
    }

    if hand.Split() != nil {
        t.Errorf("Split a hand when there was only one card left in the deck")
    }
}

func TestDiscardHand(t *testing.T) {
    deck := cards.NewDeck()
    var hand *blackjack.Hand

    for i := 0; i < 26; i++ {
        hand = blackjack.NewHand(deck)
    }

    if (blackjack.NewHand(deck) != nil) {
        t.Errorf("Deck has more than 26 blackjack hands")
    }

    hand.Discard()

    if (blackjack.NewHand(deck) == nil) {
        t.Errorf("Couldn't deal a hand after discarding one")
    }

    if (blackjack.NewHand(deck) != nil) {
        t.Errorf("Able to deal a hand after dealing 26, discarding and redealing one")
    }
}
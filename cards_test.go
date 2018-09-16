package main

import "./cards"
import "testing"
import "fmt"

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
    var hand[52] *cards.Card

    for i := 0; i < 52; i++ {
        hand[i] = deck.DealCard()

        if !ranks[hand[i].Rank] {
            t.Errorf("Unexpected rank %s", hand[i].Rank)
        }

        if !suits[hand[i].Suit] {
            t.Errorf("Unexpected suit %s", hand[i].Suit)
        }
    }

    if deck.DealCard() != nil {
        t.Errorf("Expected nil after dealing all 52 cards")
    }

    for i := 0; i < 51; i++ {
        for j := i + 1; j < 52; j++ {
            if hand[i].Rank == hand[j].Rank && hand[i].Suit == hand[j].Suit {
                t.Errorf("Card %d (%s/%s) is the same as card %d (%s/%s)",
                         i, hand[i].Rank, hand[i].Suit, j, hand[j].Rank, hand[j].Suit)
            }
        }
    }

    deck2 := cards.NewDeck()
    deck3 := cards.NewDeck()
    diffs := 0

    for i := 0; i < 52; i++ {
        card2 := deck2.DealCard()
        card3 := deck3.DealCard()

        if card2.Rank != card3.Rank || card2.Suit != card3.Suit {
            diffs++
        } else {
            fmt.Printf("Card %d is %s/%s in both decks\n", i, card2.Rank, card2.Suit)
        }
    }

    if diffs == 0 {
        t.Errorf("Two decks are identical");
    }
}

func TestPile(t *testing.T) {
    pile := cards.NewPile()

    if pile.ToStr() != "" {
        t.Errorf("Empty pile's string is '%s', not ''", pile.ToStr())
    }

    if pile.NumCards() != 0 {
        t.Errorf("Empty pile has %d cards", pile.NumCards())
    }
}

func TestDeckExhaustion(t *testing.T) {
    deck := cards.NewDeck()

    for i := 0; i < 52; i++ {
        deck.DealCard()
    }

    if (deck.DealCard() != nil) {
        t.Errorf("Deck has more than 52 cards")
    }
}

func TestDiscard(t *testing.T) {
    deck := cards.NewDeck()

    for i := 0; i < 52; i++ {
        deck.Discard(deck.DealCard())
    }

    for i := 0; i < 52; i++ {
        if deck.DealCard() == nil {
            t.Errorf("Failed to deal card %d after shuffling", i)
            break
        }
    }

    if (deck.DealCard() != nil) {
        t.Errorf("Redealt deck has more than 52 cards")
    }
}
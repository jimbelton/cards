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

    for i := 0; i < 51; i++ {
        for j := i + 1; j < 52; j++ {
            if hand[i].Rank == hand[j].Rank && hand[i].Suit == hand[j].Suit {
                t.Errorf("Card %d (%s/%s) is the same as card %d (%s/%s)",
                         i, hand[i].Rank, hand[i].Suit, j, hand[j].Rank, hand[j].Suit)
            }
        }
    }
}

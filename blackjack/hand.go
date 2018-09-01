package blackjack

import "../cards"

type Hand struct {
    cards [] *cards.Card
}

func NewHand(deck *cards.Deck) *Hand {
    var hand *Hand = new(Hand)

    hand.cards = append(hand.cards, deck.DealCard())
    hand.cards = append(hand.cards, deck.DealCard())
    return hand
}

package blackjack

import "../cards"

type Hand struct {
    cards.Pile
}

func NewHand(deck *cards.Deck) *Hand {
    var hand *Hand = new(Hand)

    hand.PutDown(deck.DealCard())
    hand.PutDown(deck.DealCard())
    return hand
}

func (hand *Hand) Score() int {
    return 4
}

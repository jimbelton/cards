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

func (hand *Hand) Score() int {
    return 4
}

func (hand *Hand) ToStr() string {
    str := ""

    for i := 0; i < len(hand.cards); i++ {
        str += hand.cards[i].Rank + "/" + hand.cards[i].Suit

        if i < len(hand.cards) - 1 {
            str += ", "
        }
    }

    return str
}
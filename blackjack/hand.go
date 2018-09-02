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
    score := 0

    for i := 0; i < len(hand.Cards); i++ {
        switch hand.Cards[i].Rank {
        case "A":  score += 1
        case "2":  score += 2
        case "3":  score += 3
        case "4":  score += 4
        case "5":  score += 5
        case "6":  score += 6
        case "7":  score += 7
        case "8":  score += 8
        case "9":  score += 9
        case "10": score += 10
        case "J":  score += 10
        case "Q":  score += 10
        case "K":  score += 10
        }
    }

    return score
}

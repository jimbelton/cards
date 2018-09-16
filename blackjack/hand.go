package blackjack

import "../cards"

type Hand struct {
    cards.Pile
    deck *cards.Deck
}

func NewHand(deck *cards.Deck) *Hand {
    var hand *Hand = new(Hand)

    hand.deck = deck

    for i := 0; i < 2; i++ {
        if (hand.PutDown(hand.deck.DealCard()) == nil) {
            return nil
        }
    }

    return hand
}

func (hand *Hand) ScoreCard(i int) int {
    switch hand.Cards[i].Rank {
    case "A": return 11
    case "2": return 2
    case "3": return 3
    case "4": return 4
    case "5": return 5
    case "6": return 6
    case "7": return 7
    case "8": return 8
    case "9": return 9
    }

    return 10
}

func (hand *Hand) Score() int {
    score := 0
    aces  := 0

    for i := 0; i < len(hand.Cards); i++ {
        score += hand.ScoreCard(i)

        if hand.Cards[i].Rank == "A" {
            aces++
        }
    }

    for score > 21 && aces > 0 {
        score -= 10
        aces--
    }

    return score
}

func (hand *Hand) IsBlackjack() bool {
    if hand.Score() == 21 && hand.NumCards() == 2 {
        return true
    }

    return false
}

func (hand *Hand) CanSplit() bool {
    if len(hand.Cards) == 2 && hand.Cards[0].Rank == hand.Cards[1].Rank {
        return true
    }

    return false
}

func (hand *Hand) Split() *Hand {
    if !hand.CanSplit() {
        return nil
    }

    var split *Hand = new(Hand)
    split.deck = hand.deck
    split.PutDown(hand.PickUp())

    if hand.PutDown(hand.deck.DealCard()) == nil ||  split.PutDown(hand.deck.DealCard()) == nil {
        return nil
    }

    return split
}

func (hand *Hand) Hit() *cards.Card {
    if hand.PutDown(hand.deck.DealCard()) != nil {
        return hand.Cards[len(hand.Cards) - 1]
    }

    return nil
}

func (hand *Hand) Discard() {
    for card := hand.PickUp(); card != nil; card = hand.PickUp() {
        hand.deck.Discard(card)
    }
}

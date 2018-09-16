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

func (hand *Hand) Score() int {
    score := 0
    aces  := 0

    for i := 0; i < len(hand.Cards); i++ {
        switch hand.Cards[i].Rank {
        case "A":  aces++
                   score += 11
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

func (hand *Hand) Hit() *Hand {
    if hand.PutDown(hand.deck.DealCard()) != nil {
        return hand
    }

    return nil
}

func (hand *Hand) Discard() {
    for card := hand.PickUp(); card != nil; card = hand.PickUp() {
        hand.deck.Discard(card)
    }
}

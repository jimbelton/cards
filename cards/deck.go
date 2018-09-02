package cards

import "math/rand"

var ranks = []string {"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var suits = []string {"hearts", "spades", "diamonds", "clubs"}

type Deck struct {
    Pile
}

func NewDeck() *Deck {
    var deck *Deck = new(Deck)

    for rank := 0; rank < len(ranks); rank++ {
        for suit := 0; suit < len(suits); suit++ {
            deck.PutDown(NewCard(ranks[rank], suits[suit]))
        }
    }

    deck.Shuffle()
    return deck
}

func (deck *Deck) DealCard() *Card {
    if len(deck.Cards) == 0 {
        return nil
    }

    card := deck.Cards[len(deck.Cards) - 1]
    deck.Cards = deck.Cards[:len(deck.Cards) - 1]
    return card
}

func (deck *Deck) Shuffle() {
    maxIndex := len(deck.Cards) - 1

    for i := 0; i < maxIndex; i++ {
        save         := deck.Cards[i]
        j            := rand.Intn(maxIndex - i) + i + 1
        deck.Cards[i] = deck.Cards[j]
        deck.Cards[j] = save
    }
}

func (deck *Deck) Stack(card *Card) *Deck {
    deck.PutDown(card)
    return deck
}

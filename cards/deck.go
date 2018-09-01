package cards

import "math/rand"

var ranks = []string {"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var suits = []string {"hearts", "spades", "diamonds", "clubs"}

type Deck struct {
    cards [52] *Card
    count int
}

func NewDeck() *Deck {
    var deck *Deck = new(Deck)

    for rank := 0; rank < len(ranks); rank++ {
        for suit := 0; suit < len(suits); suit++ {
            deck.cards[suit * len(ranks) + rank] = NewCard(ranks[rank], suits[suit])
        }
    }

    deck.count = len(suits) * len(ranks)
    deck.Shuffle()
    return deck
}

func (deck *Deck) DealCard() *Card {
    if deck.count == 0 {
        return nil
    }

    deck.count--
    return deck.cards[deck.count]
}

func (deck *Deck) Shuffle() {
    for i := 0; i < deck.count - 1; i++ {
        save         := deck.cards[i]
        j            := rand.Intn(deck.count - i - 1) + i + 1
        deck.cards[i] = deck.cards[j]
        deck.cards[j] = save
    }
}
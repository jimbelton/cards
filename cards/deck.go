package cards

var ranks = []string {"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var suits = []string {"hearts", "spades", "diamonds", "clubs"}

type Deck struct {
    cards [52] *Card
    count uint
}

func NewDeck() *Deck {
    var deck *Deck = new(Deck)

    for rank := 0; rank < len(ranks); rank++ {
        for suit := 0; suit < len(suits); suit++ {
            deck.cards[suit * len(ranks) + rank] = NewCard(ranks[rank], suits[suit])
        }
    }

    deck.count = uint(len(suits) * len(ranks))
    return deck
}

func (deck *Deck) DealCard() *Card {
    if deck.count == 0 {
        return nil
    }

    deck.count--
    return deck.cards[deck.count]
}
package cards

type Deck struct {
}

func NewDeck() *Deck {
    return &Deck{}
}

func (deck Deck) DealCard() *Card {
    return NewCard()
}
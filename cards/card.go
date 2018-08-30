package cards

type Card struct {
     Rank string
     Suit string
}

func NewCard(rank string, suit string) *Card {
    return &Card {rank, suit}
}

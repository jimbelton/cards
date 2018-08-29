package cards

type Card struct {
     Rank string
     Suit string
}

func NewCard() *Card {
    return &Card {"A", "hearts"}
}
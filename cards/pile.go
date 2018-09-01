package cards

type Pile struct {
    cards [] *Card
}

func NewPile() *Pile {
    return new(Pile)
}

func (pile *Pile) ToStr() string {
    str := ""

    for i := 0; i < len(pile.cards); i++ {
        str += pile.cards[i].Rank + "/" + pile.cards[i].Suit

        if i < len(pile.cards) - 1 {
            str += ", "
        }
    }

    return str
}

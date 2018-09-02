package cards

type Pile struct {
    Cards [] *Card
}

func NewPile() *Pile {
    return new(Pile)
}

func (pile *Pile) ToStr() string {
    str := ""

    for i := 0; i < len(pile.Cards); i++ {
        str += pile.Cards[i].Rank + "/" + pile.Cards[i].Suit

        if i < len(pile.Cards) - 1 {
            str += ", "
        }
    }

    return str
}

func (pile *Pile) PutDown(card *Card) *Pile {
    pile.Cards = append(pile.Cards, card)
    return pile
}
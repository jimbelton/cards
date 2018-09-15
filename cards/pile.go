package cards

type Pile struct {
    Cards [] *Card
}

func NewPile() *Pile {
    return new(Pile)
}

func (pile *Pile) NumCards() int {
    return len(pile.Cards)
}

func (pile *Pile) ToStr() string {
    str := ""

    for i := 0; i < len(pile.Cards); i++ {
        str += pile.Cards[i].ToStr()

        if i < len(pile.Cards) - 1 {
            str += ", "
        }
    }

    return str
}

func (pile *Pile) PutDown(card *Card) *Pile {
    if card == nil {
        return nil
    }

    pile.Cards = append(pile.Cards, card)
    return pile
}

func (pile *Pile) PickUp() *Card {
    if len(pile.Cards) == 0 {
        return nil
    }

    card := pile.Cards[len(pile.Cards) - 1]
    pile.Cards = pile.Cards[:len(pile.Cards) - 1]
    return card
}

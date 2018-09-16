package main

import "fmt"
import "./cards"
import "./blackjack"

func getYesOrNo(prompt string) bool {
    var input string

    for true {
        fmt.Print(prompt)
        fmt.Scanln(&input)

        if input[0] == 'y' {
            return true
        } else if input[0] == 'n' {
            return false
        }

        fmt.Print("Please answer 'y' or 'n'\n");
    }

    return false
}

func playHand(dealersHand *blackjack.Hand, playersHand *blackjack.Hand) {
    for playersHand.Score() <= 21 {
        hit := getYesOrNo("Do you want to hit [y/n]? ")

        if !hit {
            break
        }

        card := playersHand.Hit()
        fmt.Printf("You were dealt the %s\n", card.ToStr())
    }

    if playersHand.Score() > 21 {
        fmt.Printf("You went bust\n")
        return
    }
}

func main() {
    deck := cards.NewDeck()
    play := getYesOrNo("Do you want to play blackjack [y/n]? ")

    for play {
        dealersHand := blackjack.NewHand(deck)
        fmt.Printf("Dealer's top card is the %s.\n", dealersHand.Cards[0].ToStr())
        playersHand := blackjack.NewHand(deck)
        fmt.Printf("Your hand is %s.\n", playersHand.ToStr())

        if playersHand.IsBlackjack() {
            fmt.Print("Blackjack!\n")

            if dealersHand.ScoreCard(0) >= 10 {
                fmt.Printf("Dealer's bottom card is the %s\n", dealersHand.Cards[1].ToStr())
            }

            if dealersHand.IsBlackjack() {
                fmt.Print("Blackjack!\n")
                fmt.Print("The game is a push!\n")
            } else {
                fmt.Print("You win!\n")
            }
        } else {
            playHand(dealersHand, playersHand)
        }

        dealersHand.Discard()
        playersHand.Discard()
        play = getYesOrNo("Do you want to play again [y/n]? ")
    }
}
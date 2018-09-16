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

func main() {
    deck := cards.NewDeck()
    play := getYesOrNo("Do you want to play blackjack [y/n]? ")

    for play {
        dealersHand := blackjack.NewHand(deck)
        fmt.Printf("Dealer's top card is the %s.\n", dealersHand.Cards[0].ToStr())
        playersHand := blackjack.NewHand(deck)
        fmt.Printf("Your hand is %s.\n", playersHand.ToStr())

        if playersHand.IsBlackjack() {
            fmt.Printf("Blackjack!\n")
            fmt.Printf("You win!\n")
        }

        dealersHand.Discard()
        playersHand.Discard()
        play = getYesOrNo("Do you want to play again [y/n]? ")
    }
}
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
        play = getYesOrNo("Do you want to play again [y/n]? ")
    }
}
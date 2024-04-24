package main

import (
	"fmt"
	"os"
)

type Card struct {
	Suit string
	Face string
}

type Deck map[string](Card)

func newDeck() (*Deck) {
	suits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	faces := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	deck := Deck{}

	for _, suit := range suits {
		for _, face := range faces {
			deck[fmt.Sprintf("%s of %s", face, suit)] = Card{
				Suit: suit,
				Face: face,
			}
		}
	}

	return &deck
}

func (d Deck) print() {
	for name, _ := range d {
		fmt.Println(name)
	}
}

// func listen() {

// }

// func handleCommand() {

// }

func deck() {
	var answer string
	fmt.Print("Type an option:\nNew Deck\nPrint Current\nExit")

	for {	
		fmt.Scanln(answer)  
		

		switch answer {

			case "New Deck":
				deck := newDeck()
				deck.print()
			
			case "Print Current":
			
			case "Exit":
				os.Exit(0)
		}
	}	
}

func main() {
	// if os.Args[1] == "deck"{
	// 	deck()
	// }

	deck()
}
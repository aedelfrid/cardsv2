package cmd

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

func Cards() {
	var (
		answer int
		deck Deck
	)
	fmt.Print("Type an option:\n1. New Deck\n2. Print Current\n3. Exit\n")

	for {	
		fmt.Scan(&answer)  
		

		switch answer {

			case 1:
				deck := newDeck()
				deck.print()
			
			case 2:
				deck.print()
			case 3:
				os.Exit(0)
		}
	}	
}
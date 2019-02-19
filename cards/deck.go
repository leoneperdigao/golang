package main

import "fmt"

//Create a new typ of 'deck'
//wich is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handsize int) {

}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

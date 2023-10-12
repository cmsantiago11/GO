package main

import "fmt"

// Crate a new type of 'deck'
// which is a slice of strings

type deck []string // deck es un slice de strings

func newDeck() deck { // newDeck es una funcion que retorna un slice de strings
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit) // append es una funcion que agrega un elemento al slice
		}
	}
	return cards
}

func (d deck) print() {
	for i, carta := range d {
		fmt.Println(i, carta)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

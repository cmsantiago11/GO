package main

import "fmt"

func main() {
	//var card string = "Ace of Spades" es una manera de nombrar la variable
	card := newCard() // es otra manera de nombrar la variable igual de valida el ":=" es para declarar una variable NUEVA!
	//card = "Five of Diamonds"  de esta manera se puede cambiar el valor de la variable

	cards := []string{card, "Ace of diamonds"} // es un slice de strings
	cards = append(cards, "Six of Spades")     // append es una funcion que agrega un elemento al slice

	for i, carta := range cards { // el for range es para iterar sobre un slice
		fmt.Println(i, carta)
	}
}

func newCard() string {
	return "Five of Diamonds"
}

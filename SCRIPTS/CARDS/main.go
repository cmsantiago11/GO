package main

func main() {
	//var card string = "Ace of Spades" es una manera de nombrar la variable
	cards := newDeck() // es otra manera de nombrar la variable igual de valida el ":=" es para declarar una variable NUEVA!
	//card = "Five of Diamonds"  de esta manera se puede cambiar el valor de la variable
	hand, cardsRemaining := deal(cards, 5)
	hand.print()
	cardsRemaining.print()
}

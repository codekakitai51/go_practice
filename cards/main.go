package main

//shuffle and print
func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

// //save to file
// func main() {
// 	cards := newDeck()
// 	cards.saveToFile("cards")
// }

//make new deck from file
// func main() {
// 	cards := newDeckFromFile("cards")
// 	cards.print()
// }


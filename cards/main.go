package main //main packages are executable. All others are reusable type of packages.

import (
	"fmt" //imports other packages. Here 'fmt' package is imported.
)

func main() {
	//var card string ="Ace of Spades"  First way of variable decalaration

	//card := "Ace of Spades"   // Second way of variable declaration

	//card = "Five of Diamonds"  Value assigning after declaration

	card := newCard()

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Ace of Spades")

	for i, tempCard := range cards {
		fmt.Println(i, tempCard)
	}
	fmt.Println(card)

}

func newCard() string {
	return "Five of diamonds"
}

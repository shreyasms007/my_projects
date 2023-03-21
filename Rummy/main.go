package main

import "fmt"

func main() {

	cards := newDeck()

	cards.print()

}

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"spades", "diamonds", "hearts", "clubs"}
	cardValues := []string{"Ace", "two", "three", "four"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {

			cards = append(cards, values+" of "+suits)

		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

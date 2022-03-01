package main

import "fmt"

func main() {
	card := newDeck()
	hsize, hand := deal(card, 10)
	fmt.Println("size ", hsize, "\n\nhand ", hand)
	//card.shuffe()
	//card.print()
	
}


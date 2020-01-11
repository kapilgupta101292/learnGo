package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_cards1")
	fmt.Println(cards.toString())
}

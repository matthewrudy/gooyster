package main

import "fmt"

func main() {
	fmt.Println("OysterCard simulation!!!!!\n")

	card := NewCard()
	log(card, "Fresh new card")

	card.Credit(3000)
	log(card, "Topped up £30")

}

func log(card Card, action string) {
	fmt.Println("Action:", action)
	fmt.Println("Balance:", card.Balance())
	fmt.Println()
}

package main

import "fmt"

func main() {
	fmt.Printf("OysterCard simulation!!!!!\n\n")

	card := NewCard()
	log(card, "Fresh new card")

	card.Credit(3000)
	log(card, "Topped up £30")

}

func log(card Card, action string) {
	fmt.Println("Action:", action)
	fmt.Printf("Balance: £%.2f\n\n", float64(card.Balance())*0.01)
}

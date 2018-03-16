package main

import "fmt"

func main() {
	fmt.Printf("OysterCard simulation!!!!!\n\n")

	card := NewCard()
	log(card, "Fresh new card")

	card.TopUp(3000)
	log(card, "Topped up £30")

	card.Enter(Tube, Holborn)
	log(card, "Entered tube at Holborn")

	card.Exit(Tube, EarlsCourt)
	log(card, "Exited tube at Earl's Court")

	card.Enter(Bus, EarlsCourt)
	log(card, "Got on Bus at Earl's Court")

	card.Exit(Bus, Chelsea)
	log(card, "Got off Bus at Chelsea")

	card.Enter(Tube, EarlsCourt)
	log(card, "Entered tube at Earl's Court")

	card.Exit(Tube, Hammersmith)
	log(card, "Exited tube at Hammersmith")
}

func log(card Card, action string) {
	fmt.Println("Action:", action)
	fmt.Printf("Balance: £%.2f\n\n", float64(card.Balance())*0.01)
}

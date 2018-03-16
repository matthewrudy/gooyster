package main

import (
	"testing"
)

func TestCard_Balance(t *testing.T) {
	card := NewCard()
	assertBalance(t, card, 0)

	card.Credit(3000)
	assertBalance(t, card, 3000)

	card.Debit(350)
	assertBalance(t, card, 2650)
}

func TestCard_Enter_and_Exit(t *testing.T) {
	card := NewCard()
	card.Credit(3000)
	assertBalance(t, card, 3000)

	card.Enter(Tube, Holborn)
	assertBalance(t, card, 2680)

	card.Exit(Tube, EarlsCourt)
	assertBalance(t, card, 2750)

	card.Enter(Bus, EarlsCourt)
	assertBalance(t, card, 2570)

	card.Exit(Bus, Chelsea)
	assertBalance(t, card, 2570)

	card.Enter(Tube, EarlsCourt)
	assertBalance(t, card, 2250)

	card.Exit(Tube, Hammersmith)
	assertBalance(t, card, 2370)
}

func assertBalance(t *testing.T, card Card, expected int) {
	balance := card.Balance()
	if balance != expected {
		t.Errorf("card should have balance of %d, actual: %d", expected, balance)
	}
}

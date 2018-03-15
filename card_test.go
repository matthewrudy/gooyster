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

func assertBalance(t *testing.T, card Card, expected int) {
	balance := card.Balance()
	if balance != expected {
		t.Errorf("card should have balance of %d, actual: %d", expected, balance)
	}
}

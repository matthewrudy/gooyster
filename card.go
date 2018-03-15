package main

type Card struct {
	transactions []Transaction
}

type Transaction struct {
	amount int // GBP amount in pence
}

func NewCard() Card {
	return Card{}
}

func (c *Card) Credit(amount int) {
	c.transactions = append(c.transactions, Transaction{amount})
}

func (c *Card) Debit(amount int) {
	c.transactions = append(c.transactions, Transaction{-amount})
}

func (c *Card) Balance() int {
	sum := 0
	for _, t := range c.transactions {
		sum += t.amount
	}
	return sum
}

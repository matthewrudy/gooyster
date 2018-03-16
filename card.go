package main

type Card struct {
	transactions   []Transaction
	currentJourney Journey
}

type Transaction struct {
	amount int // GBP amount in pence
}

func NewCard() Card {
	return Card{}
}

func (c *Card) TopUp(amount int) {
	c.transactions = append(c.transactions, Transaction{amount})
}

func (c *Card) Enter(transport Transport, station *Station) {
	c.currentJourney = NewJourney(transport, station)
	c.transactions = append(c.transactions, Transaction{-c.currentJourney.Fare()})
}

func (c *Card) Exit(transport Transport, station *Station) {
	// TODO: handle unexpected state
	journey := c.currentJourney
	journey.Complete(station)
	activeTransaction := &c.transactions[len(c.transactions)-1]
	activeTransaction.amount = -journey.Fare()
}

func (c *Card) Balance() int {
	sum := 0
	for _, t := range c.transactions {
		sum += t.amount
	}
	return sum
}

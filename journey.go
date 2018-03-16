package main

type Journey struct {
	transport Transport
	entry     *Station
	exit      *Station
}

func NewJourney(transport Transport, entry *Station) Journey {
	return Journey{
		transport: transport,
		entry:     entry,
		exit:      Unknown,
	}
}

func (j *Journey) Fare() int {
	// Bus is always 180
	if j.transport == Bus {
		return 180
	}

	return 320
}

func (j *Journey) Complete(exit *Station) {
	j.exit = exit
}

type Transport int

const (
	Bus Transport = iota
	Tube
)

type Station struct {
	name  string
	zones []int
}

// a placeholder for an unknown station
// added as exit for all new journeys
var Unknown = &Station{
	name:  "Unknown",
	zones: []int{},
}

var Holborn = &Station{
	name:  "Holborn",
	zones: []int{1},
}

var EarlsCourt = &Station{
	name:  "Earl's Court",
	zones: []int{1, 2},
}

var Wimbledon = &Station{
	name:  "Wimbledon",
	zones: []int{3},
}

var Hammersmith = &Station{
	name:  "Hammersmith",
	zones: []int{2},
}

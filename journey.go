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

	// until we know the exit, we charge the max: 320
	if j.exit == Unknown {
		return 320
	}

	zoneDiff := 2 // max is 3 zones

	for _, z1 := range j.entry.zones {
		for _, z2 := range j.exit.zones {
			diff := z2 - z1
			if diff < 0 {
				diff = -diff
			}
			if diff < zoneDiff {
				zoneDiff = diff
			}
		}
	}

	isZone1 := false

	if len(j.entry.zones) == 1 && j.entry.zones[0] == 1 {
		isZone1 = true
	}

	if len(j.exit.zones) == 1 && j.exit.zones[0] == 1 {
		isZone1 = true
	}

	if zoneDiff == 0 {
		if isZone1 {
			return 250
		}
		return 200
	}

	if zoneDiff == 1 {
		if isZone1 {
			return 300
		}
		return 225
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

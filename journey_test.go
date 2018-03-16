package main

import "testing"

func TestJourney_Fare_BusIsAlways_180(t *testing.T) {
	journey := NewJourney(Bus, Holborn)
	assertFare(t, journey, 180)

	journey.Complete(Wimbledon)
	assertFare(t, journey, 180)
}

func TestJourney_Fare_StartsAtMaximum(t *testing.T) {
	zone1 := NewJourney(Tube, Holborn)
	assertFare(t, zone1, 320)

	zone3 := NewJourney(Tube, Wimbledon)
	assertFare(t, zone3, 320)
}

func TestJourney_Fare_Zone1(t *testing.T) {
	journey := NewJourney(Tube, Holborn)
	assertFare(t, journey, 320)

	journey.Complete(EarlsCourt)
	assertFare(t, journey, 250)
}

func TestJourney_Fare_Zone2(t *testing.T) {
	journey := NewJourney(Tube, EarlsCourt)
	assertFare(t, journey, 320)

	journey.Complete(Hammersmith)
	assertFare(t, journey, 200)
}

func TestJourney_Fare_Zone3(t *testing.T) {
	journey := NewJourney(Tube, Wimbledon)
	assertFare(t, journey, 320)

	journey.Complete(Wimbledon)
	assertFare(t, journey, 200)
}

func TestJourney_Fare_Zone1To2(t *testing.T) {
	journey := NewJourney(Tube, Holborn)
	assertFare(t, journey, 320)

	journey.Complete(Hammersmith)
	assertFare(t, journey, 300)
}

func TestJourney_Fare_Zone2To3(t *testing.T) {
	journey := NewJourney(Tube, Hammersmith)
	assertFare(t, journey, 320)

	journey.Complete(Wimbledon)
	assertFare(t, journey, 225)
}

func TestJourney_Fare_Zone1To3(t *testing.T) {
	journey := NewJourney(Tube, Holborn)
	assertFare(t, journey, 320)

	journey.Complete(Wimbledon)
	assertFare(t, journey, 320)
}

func assertFare(t *testing.T, journey Journey, expected int) {
	fare := journey.Fare()
	if fare != expected {
		t.Errorf("fare from %s to %s should be %d, actual: %d", journey.entry.name, journey.exit.name, expected, fare)
	}
}

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

func assertFare(t *testing.T, journey Journey, expected int) {
	fare := journey.Fare()
	if fare != expected {
		t.Errorf("fare should be %d, actual: %d", expected, fare)
	}
}

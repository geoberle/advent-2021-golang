package main

import (
	"testing"
)

func TestHotSpotsHV(t *testing.T) {
	hotspots := CalcHV(ReadInputIntoChannel("testinput.txt"))
	if hotspots != 5 {
		t.Errorf("Nr of hot spots wrong, got: %d, want: %d.", hotspots, 5)
	}
}

func TestHotSpotsHVD(t *testing.T) {
	hotspots := CalcHVD(ReadInputIntoChannel("testinput.txt"))
	if hotspots != 12 {
		t.Errorf("Nr of hot spots wrong, got: %d, want: %d.", hotspots, 12)
	}
}

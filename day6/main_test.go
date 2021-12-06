package main

import (
	"testing"
)

func TestParseInput(t *testing.T) {
	var expected = []int64{0, 1, 1, 2, 1, 0, 0, 0, 0}
	var parsed = ParseInput("testinput.txt")
	for i, e := range expected {
		if e != parsed[i] {
			t.Errorf("Population of age %d wrong, got: %d, want: %d.", i, parsed[i], e)
		}
	}
}

func Test18Days(t *testing.T) {
	phases := ParseInput("testinput.txt")
	population := SimulatePopulation(phases, 18)
	if population != 26 {
		t.Errorf("Population wrong, got: %d, want: %d.", population, 26)
	}
}

func Test18DaysOffset(t *testing.T) {
	phases := ParseInput("testinput.txt")
	population := SimulatePopulationOffset(phases, 18)
	if population != 26 {
		t.Errorf("Population wrong, got: %d, want: %d.", population, 26)
	}
}

/*func Test80Days(t *testing.T) {
	phases := ParseInput("testinput.txt")
	population := SimulatePopulation(phases, 80)
	if population != 5934 {
		t.Errorf("Population wrong, got: %d, want: %d.", population, 5934)
	}
}

func Test256Days(t *testing.T) {
	phases := ParseInput("testinput.txt")
	population := SimulatePopulation(phases, 256)
	if population != 26984457539 {
		t.Errorf("Population wrong, got: %d, want: %d.", population, 26984457539)
	}
}
*/

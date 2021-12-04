package main

import (
	"testing"
)

func TestProcessSonar(t *testing.T) {
	increases := ProcessSonar([]int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	if increases != 7 {
		t.Errorf("Wrong number of increases detected, got: %d, want: %d.", increases, 7)
	}
}

func TestProcessSectionSonar(t *testing.T) {
	increases := ProcessSectionSonar([]int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	if increases != 5 {
		t.Errorf("Wrong number of increase sections detected, got: %d, want: %d.", increases, 5)
	}
}

func TestProcessSectionSonarLineByLine(t *testing.T) {
	increases := ProcessSectionSonarLineByLine("testinput.txt")
	if increases != 5 {
		t.Errorf("Wrong number of increase sections detected, got: %d, want: %d.", increases, 5)
	}
}

func TestProcessSectionSonarFromChannel(t *testing.T) {
	c := ReadInputIntoChannel("testinput.txt")
	increases := ProcessSectionSonarFromChannel(c)
	if increases != 5 {
		t.Errorf("Wrong number of increase sections detected, got: %d, want: %d.", increases, 5)
	}
}

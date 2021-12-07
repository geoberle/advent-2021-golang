package day7

import (
	"fmt"
	"testing"
)

func TestMinFuelTestInput(t *testing.T) {
	input := []int64{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	fuel := CalcMinFuelPosition(input)
	if fuel != 37 {
		t.Errorf("Fuel wrong, got: %d, want: %d.", fuel, 37)
	}
}

func TestMinFuel(t *testing.T) {
	input := ParseInput("input.txt")
	fuel := CalcMinFuelPosition(input)
	if fuel != 355521 {
		t.Errorf("Fuel wrong, got: %d, want: %d.", fuel, 355521)
	}
	fmt.Printf("Part 1 - %d", fuel)
}

func TestFuelIncAt5(t *testing.T) {
	input := []int64{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	fuel := CalcFuelIncAtPosition(input, 5)
	if fuel != 168 {
		t.Errorf("Fuel wrong, got: %d, want: %d.", fuel, 168)
	}
}

func TestMinFuelInc(t *testing.T) {
	input := ParseInput("input.txt")
	fuel := CalcMinFuelInc(input)
	if fuel != 100148777 {
		t.Errorf("Fuel wrong, got: %d, want: %d.", fuel, 100148777)
	}
	fmt.Printf("Part 2 - %d", fuel)
}

package main

import (
	"testing"
)

func TestGammaEpsilon(t *testing.T) {
	c := ReadInputIntoChannel("testinput.txt")
	g, e := Rates(c)
	product := g * e
	if product != 198 {
		t.Errorf("Gamma*Epsilon wrong, got: %d, want: %d.", product, 198)
	}
}
func TestOxygen(t *testing.T) {
	oxygen := LifeSupportMetric(ReadInput("testinput.txt"), "oxygen")
	if oxygen != 23 {
		t.Errorf("Oxygen wrong, got %d, want %d.", oxygen, 23)
	}
}

func TestCO2(t *testing.T) {
	co2 := LifeSupportMetric(ReadInput("testinput.txt"), "co2")
	if co2 != 10 {
		t.Errorf("CO2 wrong, got %d, want %d.", co2, 10)
	}
}

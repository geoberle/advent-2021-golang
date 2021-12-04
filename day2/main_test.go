package main

import (
	"testing"
)

func TestNavigate(t *testing.T) {
	c := ReadInputIntoChannel("testinput.txt")
	position := Navigate(c)
	product := position.horizontal * position.depth
	if product != 150 {
		t.Errorf("Product of horizontal position and depth wrong, got: %d, want: %d.", product, 150)
	}
}

func TestAim(t *testing.T) {
	c := ReadInputIntoChannel("testinput.txt")
	position := Aim(c)
	product := position.horizontal * position.depth
	if product != 900 {
		t.Errorf("Product of horizontal position and depth wrong, got: %d, want: %d.", product, 900)
	}
}

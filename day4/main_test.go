package main

import (
	"testing"
)

func TestWinFirst(t *testing.T) {
	turnNumbers, boards := ReadInput("testinput.txt", 5)
	win_score := FindWinningBoardGoroutine(3, turnNumbers, boards, true)
	if win_score != 4512 {
		t.Errorf("First win score wrong, got: %d, want: %d.", win_score, 4512)
	}
}

func TestWinLast(t *testing.T) {
	turnNumbers, boards := ReadInput("testinput.txt", 5)
	win_score := FindWinningBoardGoroutine(3, turnNumbers, boards, false)
	if win_score != 1924 {
		t.Errorf("Last win score wrong, got: %d, want: %d.", win_score, 1924)
	}
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type board struct {
	boardSize int
	numbers   []int64
}

func (b *board) CalcWinScore(numberTurns map[int64]int) (int, int64) {
	win_turn := math.MaxInt
	var win_number int64
	for row := 0; row < b.boardSize; row++ {
		row_win_turn := -1
		var row_win_number int64
		for col := 0; col < b.boardSize; col++ {
			field_turn, found := numberTurns[b.numbers[row*b.boardSize+col]]
			if !found {
				row_win_turn = -1
				break
			}
			if field_turn > row_win_turn {
				row_win_turn = field_turn
				row_win_number = b.numbers[row*b.boardSize+col]
			}
		}
		if row_win_turn < win_turn {
			win_turn = row_win_turn
			win_number = row_win_number
		}
	}

	for col := 0; col < b.boardSize; col++ {
		col_win_turn := -1
		var col_win_number int64
		for row := 0; row < b.boardSize; row++ {
			field_turn, found := numberTurns[b.numbers[row*b.boardSize+col]]
			if !found {
				col_win_turn = -1
				break
			}
			if field_turn > col_win_turn {
				col_win_turn = field_turn
				col_win_number = b.numbers[row*b.boardSize+col]
			}
		}
		if col_win_turn < win_turn {
			win_turn = col_win_turn
			win_number = col_win_number
		}
	}

	var unmarked_sum int64 = 0
	for row := 0; row < b.boardSize; row++ {
		for col := 0; col < b.boardSize; col++ {
			field_number := b.numbers[row*b.boardSize+col]
			field_turn, found := numberTurns[field_number]
			if !found || field_turn > win_turn {
				unmarked_sum += field_number
			}
		}
	}
	//fmt.Printf("%d %d %d\n", win_turn, win_number, unmarked_sum)
	return win_turn, win_number * unmarked_sum
}

func ReadInput(path string, boardSize int) (map[int64]int, []board) {
	file, _ := os.Open(path)
	defer file.Close()
	boards := []board{}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// make a map to track in which order numbers are pulled
	numbers := strings.Split(lines[0], ",")
	numberOrder := make(map[int64]int, len(numbers))
	for i, ns := range numbers {
		n, _ := strconv.ParseInt(ns, 10, 0)
		numberOrder[n] = i
	}

	for i := 2; i < len(lines); i += boardSize + 1 {
		boardNumers := make([]int64, boardSize*boardSize)
		for row := 0; row < boardSize; row += 1 {
			for col, ns := range strings.Fields(lines[i+row]) {
				n, _ := strconv.ParseInt(ns, 10, 0)
				boardNumers[row*boardSize+col] = n
			}
		}
		boards = append(boards, board{boardSize: boardSize, numbers: boardNumers})
	}

	return numberOrder, boards
}

type job struct {
	turnNumbers map[int64]int
	board       board
}

type jobsresult struct {
	winTurn int
	score   int64
}

func FindWinningBoardGoroutine(workers int, turnNumbers map[int64]int, boards []board, winning bool) int64 {
	jobs := make(chan job, len(boards))
	scores := make(chan jobsresult)

	// start worker
	for i := 0; i < workers; i++ {
		go func(jobs <-chan job, scores chan<- jobsresult) {
			for j := range jobs {
				turn, score := j.board.CalcWinScore(j.turnNumbers)
				scores <- jobsresult{turn, score}
			}
		}(jobs, scores)
	}

	// send work
	for _, b := range boards {
		jobs <- job{turnNumbers, b}
	}
	close(jobs)

	var win_score int64
	var win_turn int = math.MaxInt
	if !winning {
		win_turn = -1
	}
	for i := 0; i < len(boards); i++ {
		result := <-scores
		if (winning && result.winTurn < win_turn) || (!winning && result.winTurn > win_turn) {
			win_turn = result.winTurn
			win_score = result.score
		}
	}
	close(scores)

	return win_score
}

func main() {
	turnNumbers, boards := ReadInput("input.txt", 5)

	first_win_score := FindWinningBoardGoroutine(5, turnNumbers, boards, true)
	fmt.Printf("Part 1 - score = %d\n", first_win_score)

	last_win_score := FindWinningBoardGoroutine(5, turnNumbers, boards, false)
	fmt.Printf("Part 2 - score = %d\n", last_win_score)
}

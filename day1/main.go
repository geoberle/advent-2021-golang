package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func ProcessSonar(input []int64) int {
	inc := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			inc += 1
		}
	}
	return inc
}

/**
Sections overlap by two values, so the code only needs to check the last
value of the current section and the first value of the previous one.
*/
func ProcessSectionSonar(input []int64) int {
	inc := 0
	for i := 3; i < len(input); i++ {
		if input[i] > input[i-3] {
			inc += 1
		}
	}
	return inc
}

/**
While the previous solution works nicely, i don't like the idea of reading
the entire file into memory. So processing line by line and only keeping the
last 3 read numbers in a ring buffer helps with that.
*/
func ProcessSectionSonarLineByLine(path string) int {
	inc := 0
	file, _ := os.Open(path)
	defer file.Close()
	nr := 0
	ring_buffer := [3]int64{0, 0, 0}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		current, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		if nr >= 3 && current > ring_buffer[nr%3] {
			inc += 1
		}
		ring_buffer[nr%3] = current
		nr += 1
	}
	return inc
}

/*
Coming from Python i wondered, if something like a generator would
be helpful to declutter the code in ProcessSectionSonarLineByLine
while keeping only one line in memory at a time.
Channels to the rescue - read about them here: https://go.dev/blog/pipelines
The heavy lifting in the file handling is done in ReadInputIntoChannel
I'm aware that Channels are usually used a lot for parallel processing, but
they still seemed a nice way of splitting the file handling and the data
processing.
*/
func ProcessSectionSonarFromChannel(in <-chan int64) int {
	inc := 0
	nr := 0
	ring_buffer := [3]int64{0, 0, 0}
	for current := range in {
		if nr >= 3 && current > ring_buffer[nr%3] {
			inc += 1
		}
		ring_buffer[nr%3] = current
		nr += 1
	}
	return inc
}

func ReadInputIntoChannel(path string) chan int64 {
	out := make(chan int64)
	go func() {
		file, _ := os.Open(path)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			number, _ := strconv.ParseInt(scanner.Text(), 10, 0)
			out <- number
		}
		close(out)
	}()
	return out
}

func ReadInput(path string) ([]int64, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	numbers := []int64{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := strconv.ParseInt(scanner.Text(), 10, 0)
		numbers = append(numbers, number)
	}
	return numbers, nil
}

func main() {
	// part 1
	input, _ := ReadInput("input.txt")
	fmt.Printf("Part 1 increases: %d\n", ProcessSonar(input))

	// part 2
	c := ReadInputIntoChannel("input.txt")
	fmt.Printf("Part 2 increases: %d\n", ProcessSectionSonarFromChannel(c))
}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func Rates(in <-chan string) (int, int) {
	var bits []int
	line_count := 0
	for line := range in {
		line_count++
		if len(bits) == 0 {
			bits = make([]int, len(line))
		}
		for i, b := range line {
			if b == '1' {
				bits[i]++
			}
		}
	}
	gamma := 0
	for i, b := range bits {
		if b > line_count/2 {
			gamma |= 1 << (len(bits) - i - 1)
		}
	}
	epsilon := gamma ^ (int(math.Pow(float64(2), float64(len(bits)))) - 1)
	return gamma, epsilon
}

func ReadInputIntoChannel(path string) chan string {
	out := make(chan string)
	go func() {
		file, _ := os.Open(path)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			out <- strings.TrimSuffix(scanner.Text(), "\n")
		}
		close(out)
	}()
	return out
}

func LifeSupportMetric(numbers []string, metric string) int {
	var pos int
	nLen := float64(len(numbers))
	for nLen > 1 {
		var bits float64
		for _, n := range numbers {
			if n[pos] == '1' {
				bits++
			}
		}

		var filter uint8
		if metric == "oxygen" {
			filter = uint8('0')
			if bits >= nLen/2 {
				filter = '1'
			}
		} else {
			filter = uint8('0')
			if bits < nLen/2 {
				filter = '1'
			}
		}
		nLen = 0
		for _, n := range numbers {
			if n[pos] == filter {
				numbers[int(nLen)] = n
				nLen++
			}
		}
		numbers = numbers[:int(nLen)]
		pos++
	}

	rating := 0
	for i, b := range numbers[0] {
		if b == '1' {
			rating |= 1 << (len(numbers[0]) - i - 1)
		}
	}
	return rating
}

func ReadInput(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()
	numbers := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers = append(numbers, scanner.Text())
	}
	return numbers
}

func main() {
	c := ReadInputIntoChannel("input.txt")
	g, e := Rates(c)
	fmt.Printf("Part 1: gamma = %d, epsilon=%d prod=%d\n", g, e, g*e)

	oxygen := LifeSupportMetric(ReadInput("input.txt"), "oxygen")
	co2 := LifeSupportMetric(ReadInput("input.txt"), "co2")
	lifesupport_rating := oxygen * co2
	fmt.Printf("Part 2 - life support rating: %d\n", lifesupport_rating)
}

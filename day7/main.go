package day7

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"

	"github.com/montanaflynn/stats"
)

func CalcFuelAtPosition(input []int64, position int64) int64 {
	var allFuel int64 = 0
	for _, crab := range input {
		fuel := crab - position
		if fuel < 0 {
			fuel *= -1
		}
		allFuel += fuel
	}
	return allFuel
}

func CalcFuelIncAtPosition(input []int64, position int64) int64 {
	var allFuel int64 = 0
	for _, crab := range input {
		distance := crab - position
		if distance < 0 {
			distance *= -1
		}
		allFuel += distance * (distance + 1) / 2
	}
	return allFuel
}

func CalcMinFuelInc(input []int64) int64 {
	var minFuel int64 = math.MaxInt64
	for i := int64(0); i < 1500; i++ {
		fuel := CalcFuelIncAtPosition(input, i)
		if fuel < minFuel {
			minFuel = fuel
		}
	}
	return minFuel
}

func CalcMinFuelPosition(input []int64) int64 {
	data := stats.LoadRawData(input)
	median, _ := stats.Median(data)
	return CalcFuelAtPosition(input, int64(median))
}

func ParseInput(path string) []int64 {
	data, _ := ioutil.ReadFile(path)
	parts := strings.Split(strings.TrimSuffix(string(data), "\n"), ",")
	numbers := make([]int64, len(parts))
	for i, s := range parts {
		pos, _ := strconv.ParseInt(s, 10, 0)
		numbers[i] = pos
	}
	return numbers
}

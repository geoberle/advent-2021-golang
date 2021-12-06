package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func SimulatePopulation(daysToRepro []int64, daysToSimulate int) int64 {
	for day := 0; day < daysToSimulate; day++ {
		daysToRepro = []int64{daysToRepro[1], daysToRepro[2], daysToRepro[3], daysToRepro[4], daysToRepro[5], daysToRepro[6], daysToRepro[7] + daysToRepro[0], daysToRepro[8], daysToRepro[0]}
	}
	var sum int64
	for _, nrFish := range daysToRepro {
		sum += nrFish
	}
	return sum
}

// daysToRepro[0] = number of fishes that will reproduce today
// daysToRepro[1] = number of fish that will reproduce tomorrow
// ...
func SimulatePopulationOffset(daysToRepro []int64, daysToSimulate int) int64 {
	zeroIdx := 0
	for day := 0; day < daysToSimulate; day++ {
		daysToRepro[(zeroIdx+7)%9] += daysToRepro[zeroIdx]
		zeroIdx = (zeroIdx + 1) % 9
	}
	var sum int64
	for _, nrFish := range daysToRepro {
		sum += nrFish
	}
	return sum
}

func ParseInput(path string) []int64 {
	data, _ := ioutil.ReadFile(path)
	parts := strings.Split(strings.TrimSuffix(string(data), "\n"), ",")
	daysToRepro := make([]int64, 9)
	for _, s := range parts {
		days, _ := strconv.ParseInt(s, 10, 0)
		daysToRepro[days]++
	}
	return daysToRepro
}

func main() {
	phases := ParseInput("input.txt")

	population80 := SimulatePopulation(phases, 80)
	fmt.Printf("Part 1 - 80 day population = %d\n", population80)

	population256 := SimulatePopulation(phases, 256)
	fmt.Printf("Part 2 - 256 day population = %d\n", population256)

	population256offset := SimulatePopulationOffset(phases, 256)
	fmt.Printf("Part 2 - 256 day population with offset = %d\n", population256offset)
}

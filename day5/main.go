package main

import (
	"bufio"
	"fmt"
	"os"
)

func CalcHV(in <-chan string) int {
	hotspot := make(map[string]int)
	for line := range in {
		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		if x1 == x2 || y1 == y2 { // vertical and horizontal
			hotspot[fmt.Sprintf("%d-%d", x1, y1)]++
			// we move the start point towards the end point
			for x1 != x2 || y1 != y2 {
				if x1 > x2 {
					x1--
				} else if x1 < x2 {
					x1++
				}
				if y1 > y2 {
					y1--
				} else if y1 < y2 {
					y1++
				}
				hotspot[fmt.Sprintf("%d-%d", x1, y1)]++
			}
		}

	}
	var veryHotSpots int = 0
	for _, h := range hotspot {
		if h > 1 {
			veryHotSpots++
		}
	}
	return veryHotSpots
}

func CalcHVD(in <-chan string) int {
	hotspot := make(map[string]int)
	for line := range in {
		var x1, y1, x2, y2 int
		fmt.Sscanf(line, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2)
		hotspot[fmt.Sprintf("%d-%d", x1, y1)]++
		// we move the start point towards the end point
		for x1 != x2 || y1 != y2 {
			if x1 > x2 {
				x1--
			} else if x1 < x2 {
				x1++
			}
			if y1 > y2 {
				y1--
			} else if y1 < y2 {
				y1++
			}
			hotspot[fmt.Sprintf("%d-%d", x1, y1)]++
		}

	}
	var veryHotSpots int = 0
	for _, h := range hotspot {
		if h > 1 {
			veryHotSpots++
		}
	}
	return veryHotSpots
}

func ReadInputIntoChannel(path string) chan string {
	out := make(chan string)
	go func() {
		file, _ := os.Open(path)
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			out <- scanner.Text()
		}
		close(out)
	}()
	return out
}

func main() {
	// part 1
	hotspots := CalcHV(ReadInputIntoChannel("input.txt"))
	fmt.Printf("Part 1 hot spots: %d\n", hotspots)

	// part 2
	hotspotsd := CalcHVD(ReadInputIntoChannel("input.txt"))
	fmt.Printf("Part 2 hot spots incl diagonal: %d\n", hotspotsd)
}

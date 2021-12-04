package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type position struct {
	horizontal int64
	depth      int64
	aim        int64
}

/*
I used lambdas in Python to solve this challenge and i was really happy when
i found out that you can have callable function references that are type safe!!
*/
type CommandFunction func(position, int64) position

func Forward(state position, move int64) position {
	return position{
		horizontal: state.horizontal + move,
		depth:      state.depth,
		aim:        0}
}

func Up(state position, move int64) position {
	return position{
		horizontal: state.horizontal,
		depth:      state.depth - move,
		aim:        0}
}

func Down(state position, move int64) position {
	return position{
		horizontal: state.horizontal,
		depth:      state.depth + move,
		aim:        0}
}

func Navigate(in <-chan string) position {
	command_central := map[string]CommandFunction{
		"forward": Forward,
		"up":      Up,
		"down":    Down,
	}
	position := position{horizontal: 0, depth: 0, aim: 0}
	for instruction := range in {
		split := strings.Split(instruction, " ")
		command := split[0]
		arg, _ := strconv.ParseInt(split[1], 10, 0)
		position = command_central[command](position, arg)
	}
	return position
}

func ForwardAim(state position, move int64) position {
	return position{
		horizontal: state.horizontal + move,
		depth:      state.depth + (state.aim * move),
		aim:        state.aim}
}

func UpAim(state position, move int64) position {
	return position{
		horizontal: state.horizontal,
		depth:      state.depth,
		aim:        state.aim - move}
}

func DownAim(state position, move int64) position {
	return position{
		horizontal: state.horizontal,
		depth:      state.depth,
		aim:        state.aim + move}
}

func Aim(in <-chan string) position {
	command_central := map[string]CommandFunction{
		"forward": ForwardAim,
		"up":      UpAim,
		"down":    DownAim,
	}
	position := position{horizontal: 0, depth: 0, aim: 0}
	for instruction := range in {
		split := strings.Split(instruction, " ")
		command := split[0]
		arg, _ := strconv.ParseInt(split[1], 10, 0)
		position = command_central[command](position, arg)
	}
	return position
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
	c1 := ReadInputIntoChannel("input.txt")
	pos1 := Navigate(c1)
	fmt.Printf("Part 1 position: %d\n", pos1.horizontal*pos1.depth)

	// part 1
	c2 := ReadInputIntoChannel("input.txt")
	pos2 := Aim(c2)
	fmt.Printf("Part 2 position: %d\n", pos2.horizontal*pos2.depth)
}

package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	horizontal int
	vertical   int
	aim        int
}

func Day2() {
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	position := Position{
		horizontal: 0,
		vertical:   0,
		aim:        0,
	}

	for scanner.Scan() {
		position = handleInstruction(position, scanner.Text())
	}

	fmt.Println("Day 2 part 1:")
	fmt.Println(position.horizontal * position.vertical)

}

func handleInstruction(position Position, instruction string) Position {
	parts := strings.Split(instruction, " ")
	direction := parts[0]
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal("Can't convert to number, ", parts[1])
	}

	switch direction {
	case "up":
		position.aim -= value
	case "down":
		position.aim += value
	case "forward":
		position.horizontal += value
		position.vertical += position.aim * value
	}

	return position
}

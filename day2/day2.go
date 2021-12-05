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
	x int
	y int
}

func Day2() {
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	position := Position{
		x: 0,
		y: 0,
	}

	for scanner.Scan() {
		position = handleInstruction(position, scanner.Text())
	}

	fmt.Println("Day 2 part 1:")
	fmt.Println(position.x * position.y)

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
		position.y -= value
	case "down":
		position.y += value
	case "forward":
		position.x += value
	}

	return position
}

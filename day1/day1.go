package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day1() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var values []int

	for scanner.Scan() {
		intVar, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		values = append(values, intVar)
	}

	counter := 0
	for i := 1; i < len(values); i++ {
		if values[i-1] < values[i] {
			counter += 1
		}
	}
	fmt.Println(counter)
}

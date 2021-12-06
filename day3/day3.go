package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Day3() {
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var counter []int
	var length int

	for scanner.Scan() {
		binaryStr := scanner.Text()
		length += 1
		for i, bit := range binaryStr {
			if len(counter) < i+1 {
				counter = append(counter, 0)
			}
			if string(bit) == "1" {
				counter[i] = counter[i] + 1
			}
		}
	}

	gammaRate := ""
	epsilonRate := ""
	for _, counterValue := range counter {
		if counterValue > length/2 {
			gammaRate = gammaRate + "1"
			epsilonRate = epsilonRate + "0"
		} else {
			gammaRate = gammaRate + "0"
			epsilonRate = epsilonRate + "1"
		}
	}
	gammaValue, err := strconv.ParseInt(gammaRate, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilonValue, err := strconv.ParseInt(epsilonRate, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Day 3 1")
	fmt.Println(gammaValue * epsilonValue)
}

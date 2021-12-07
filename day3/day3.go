package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day3() {
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var data []string

	for scanner.Scan() {
		binaryStr := scanner.Text()
		data = append(data, binaryStr)
	}

	oxygenGeneratorRating := recursiveSearch(data, 0, false)
	co2ScrubberRating := recursiveSearch(data, 0, true)

	oxygenValue, err := strconv.ParseInt(strings.Join(oxygenGeneratorRating, ""), 2, 64)
	co2Value, err := strconv.ParseInt(strings.Join(co2ScrubberRating, ""), 2, 64)
	fmt.Println("Part 3 2")

	fmt.Println(oxygenValue * co2Value)

}

func recursiveSearch(data []string, index int, invert bool) []string {
	result := data
	binaryConditions := getBinaryConditionSringForData(result, len(result))
	if invert {
		binaryConditions = invertBinaryArray(binaryConditions)
	}
	condition := strconv.Itoa(binaryConditions[index])
	result = filterByBitAtIndex(result, index, condition)
	if len(result) > 1 {
		result = recursiveSearch(result, index+1, invert)
	}
	return result
}

func filterByBitAtIndex(data []string, index int, condition string) []string {
	var result []string
	for _, entry := range data {
		if string(entry[index]) == condition {
			result = append(result, entry)
		}
	}
	return result
}

func getBindaryConditionStringForCounter(counter []int, inputLength int) []int {
	var result []int
	for _, counterValue := range counter {
		if float64(counterValue) >= float64(inputLength)/2 {
			result = append(result, 1)
		} else {
			result = append(result, 0)
		}
	}
	return result
}

func invertBinaryArray(binaryArray []int) []int {
	var result []int
	for _, bit := range binaryArray {
		if bit == 1 {
			result = append(result, 0)
		} else {
			result = append(result, 1)
		}
	}
	return result
}

func getBinaryConditionSringForData(data []string, inputLength int) []int {
	var result []int
	for _, binaryStr := range data {
		for i, bit := range binaryStr {
			if len(result) < i+1 {
				result = append(result, 0)
			}
			if string(bit) == "1" {
				result[i] = result[i] + 1
			}
		}
	}
	return getBindaryConditionStringForCounter(result, inputLength)
}

package CalculatorKata

import (
	"strconv"
	"strings"
)

func Add(input string) int {
	numbers := getNumbers(input)
	total := 0

	for _, n := range numbers {
		total = total + n
	}

	return total
}

func getNumbers(input string) []int {
	seperated := strings.Split(input, ",")
	numberCount := len(seperated)

	result := make([]int, numberCount)

	for i := 0; i < numberCount; i++ {
		number := seperated[i]
		result[i], _ = strconv.Atoi(number)
	}

	return result
}

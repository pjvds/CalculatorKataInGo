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
	separated := strings.FieldsFunc(input, isSeparator)
	numberCount := len(separated)

	result := make([]int, numberCount)

	for i := 0; i < numberCount; i++ {
		number := separated[i]
		result[i], _ = strconv.Atoi(number)
	}

	return result
}

func isSeparator(r rune) bool {
	return r == ',' || r == '\n'
}

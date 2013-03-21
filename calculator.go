package CalculatorKata

import (
	"errors"
	"strconv"
	"strings"
)

func Add(input string) (int, error) {
	numbers := getNumbers(input)
	total := 0

	for _, n := range numbers {
		if n < 0 {
			return 0, errors.New("negatives not allowed")
		}

		total = total + n
	}

	return total, nil
}

func getSeparators(input string) ([]rune, string) {
	customSeparatorPrefix := "//"
	var separators []rune

	if strings.Index(input, customSeparatorPrefix) == 0 {
		separators = make([]rune, 1)
		separators[0] = rune(input[2])

		input = input[4:]
	} else {
		separators = make([]rune, 2)
		separators[0] = ','
		separators[1] = '\n'
	}

	return separators, input
}

func getNumbers(input string) []int {
	separators, input := getSeparators(input)

	separated := strings.FieldsFunc(input, createIsSeparator(separators))
	numberCount := len(separated)

	result := make([]int, numberCount)

	for i := 0; i < numberCount; i++ {
		number := separated[i]
		result[i], _ = strconv.Atoi(number)
	}

	return result
}

func createIsSeparator(separators []rune) func(r rune) bool {
	return func(r rune) bool {
		for _, s := range separators {
			if r == s {
				return true
			}
		}
		return false
	}
}

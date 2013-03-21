package CalculatorKata

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type input struct {
	raw        string
	delimiters []rune
	numbers    []int
}

type predicate func(item interface{}) bool

func where(slice []interface{}, p func(item interface{}) bool) []interface{} {
	result := make([]interface{}, 0)

	for _, item := range slice {
		if p(item) {
			result = append(result, item)
		}
	}

	return result
}

func Add(inputString string) (int, error) {
	input := parse(inputString)
	total := 0

	negatives := where(input.numbers, func(item int) bool {
		return item < 0
	})

	for _, n := range input.numbers {
		if n < 0 {
			msg := fmt.Sprintf("negatives not allowed: %v", n)
			return 0, errors.New(msg)
		}

		total = total + n
	}

	return total, nil
}

func parse(inputRaw string) input {
	delimiters, numberInput := getSeparators(inputRaw)
	numbers := getNumbers(numberInput, delimiters)

	return input{
		raw:        inputRaw,
		delimiters: delimiters,
		numbers:    numbers,
	}
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

func getNumbers(numberInput string, separators []rune) []int {

	separated := strings.FieldsFunc(numberInput, createIsSeparator(separators))
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

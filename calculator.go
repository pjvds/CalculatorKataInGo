package CalculatorKata

import (
	"strconv"
)

func Add(input string) int {
	if input == "" {
		return 0
	}

	number, err := strconv.Atoi(input)

	if err != nil {
		panic(err)
	}

	return number
}

package CalculatorKata

import (
	"fmt"
	"testing"
)

func TestAddEmptyStringReturnZero(t *testing.T) {
	expected := 0

	input := ""
	result := Add(input)

	if result == expected {
		t.Log("Add returns zero on empty string")
	} else {
		t.Errorf("Add returns %v on empty string, %v is expected", result, expected)
	}
}

func TestAddReturnsInputForSingleNumber(t *testing.T) {
	expected := 4

	input := fmt.Sprint(expected)
	result := Add(input)

	if result == expected {
		t.Log("Add returns %v as expected for single number input", result)
	} else {
		t.Errorf("Add returns %v on single number input, %v is expected", result, expected)
	}
}

func TestAddReturnsInputForTwoNumbers(t *testing.T) {
	expected := 4

	input := "2,2"
	result := Add(input)

	if result == expected {
		t.Log("Add returns %v as expected for two number input separated by comma", result)
	} else {
		t.Errorf("Add returns %v on two number input separated by comma, %v is expected", result, expected)
	}
}

func TestAddReturnsInputForTwoNumbersSeparatedByComma(t *testing.T) {
	expected := 4

	input := "2,2"
	result := Add(input)

	if result == expected {
		t.Log("Add returns %v as expected for two number input separated by comma", result)
	} else {
		t.Errorf("Add returns %v on two number input separated by comma, %v is expected", result, expected)
	}
}

func TestAddReturnsInputForTwoNumbersSeparatedByNewline(t *testing.T) {
	expected := 4

	input := "2\n2"
	result := Add(input)

	if result == expected {
		t.Log("Add returns %v as expected for two number input separated by newline", result)
	} else {
		t.Errorf("Add returns %v on two number input separated by newline, %v is expected", result, expected)
	}
}

func TestAddReturnsInputForNumbersSeparatedByNewlineAndComma(t *testing.T) {
	expected := 7

	input := "2,3\n2"
	result := Add(input)

	if result == expected {
		t.Log("Add returns %v as expected for numbers separated by newline and comma", result)
	} else {
		t.Errorf("Add returns %v for numbers separated by newline and comma, %v is expected", result, expected)
	}
}

func TestCustomSeparatorSupport(t *testing.T) {
	expected := 12

	input := "//;\n5;7"

	result := Add(input)

	if result != expected {
		t.Errorf("Result was %v, but %v is expected", result, expected)
	}
}

func TestGetSeparatorsSemicolon(t *testing.T) {
	input := "//;\n3;3"

	separators, input := getSeparators(input)

	expectedSize := 1
	size := len(separators)
	if size != expectedSize {
		t.Errorf("separator slice has an invalid size of %v, %v was expected", size, expectedSize)
	}

	expectedSeparator := ';'
	separator := separators[0]
	if separator != expectedSeparator {
		t.Errorf("invalid separator of %v, %v was expected", separator, expectedSeparator)
	}

	expectedInput := "3;3"
	if input != expectedInput {
		t.Errorf("invalid input result of '%v', '%v' was expected", input, expectedInput)
	}
}

func TestGetSeparatorsX(t *testing.T) {
	input := "//x\n9;12"

	separators, input := getSeparators(input)

	expectedSize := 1
	size := len(separators)
	if size != expectedSize {
		t.Errorf("separator slice has an invalid size of %v, %v was expected", size, expectedSize)
	}

	expectedSeparator := 'x'
	separator := separators[0]
	if separator != expectedSeparator {
		t.Errorf("invalid separator of %v, %v was expected", separator, expectedSeparator)
	}

	expectedInput := "9;12"
	if input != expectedInput {
		t.Errorf("invalid input result of '%v', '%v' was expected", input, expectedInput)
	}
}

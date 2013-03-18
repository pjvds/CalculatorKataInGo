# Calculator Kata in Go

## Before you start:

* Try not to read ahead.
* Do one task at a time. The trick is to learn to work incrementally.
* Make sure you only test for correct inputs. there is no need to test for invalid inputs for this kata

## String Calculator

1. Create a simple String calculator with a method int Add(string numbers)
	1. The method can take 0, 1 or 2 numbers, and will return their sum (for an empty string it will return 0) for example ââ or â1â or â1,2â
	2. Start with the simplest test case of an empty string and move to 1 and two numbers
	3. Remember to solve things as simply as possible so that you force yourself to write tests you did not think about
	4. Remember to refactor after each passing test
2. Allow the Add method to handle an unknown amount of numbers
3. Allow the Add method to handle new lines between numbers (instead of commas).
	1. the following input is ok:  â1\n2,3â  (will equal 6)
	2. the following input is NOT ok:  â1,\nâ (not need to prove it - just clarifying)
4. Support different delimiters
	1. to change a delimiter, the beginning of the string will contain a separate line that looks like this:   â//[delimiter]\n[numbersâ¦]â for example â//;\n1;2â should return three where the default delimiter is â;â .
the first line is optional. all existing scenarios should still be supported
Calling Add with a negative number will throw an exception ânegatives not allowedâ - and the negative that was passed.if there are multiple negatives, show all of them in the exception message
stop here if you are a beginner. Continue if you can finish the steps so far in less than 30 minutes.
Numbers bigger than 1000 should be ignored, so adding 2 + 1001  = 2
Delimiters can be of any length with the following format:  â//[delimiter]\nâ for example: â//[***]\n1***2***3â should return 6
Allow multiple delimiters like this:  â//[delim1][delim2]\nâ for example â//[*][%]\n1*2%3â should return 6.
make sure you can also handle multiple delimiters with length longer than one char


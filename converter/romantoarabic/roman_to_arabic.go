package romantoarabic

import (
	"errors"
	"regexp"
)

var romanToIntMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

// ToInteger Converts Roman Numeral strings to integer. Returns 0 if string is empty.
// Supports numbers between 1 and 3999. Negative numbers and lowercase letters are unsupported and will throw an error.
func ToInteger(romanNumeral string) (int, error) {
	if !IsRomanNumber(romanNumeral) {
		return 0, errors.New("given string is not a valid roman numeral")
	}

	result := 0
	ln := len(romanNumeral)
	for i := 0; i < ln; i++ {
		currentChar := string(romanNumeral[i])
		currentCharValue := romanToIntMap[currentChar]
		if i < ln-1 {
			nextChar := string(romanNumeral[i+1])
			nextCharValue := romanToIntMap[nextChar]
			if currentCharValue < nextCharValue {
				result += nextCharValue - currentCharValue
				i++
			} else {
				result += currentCharValue
			}
		} else {
			result += currentCharValue
		}
	}
	return result, nil
}

// IsRomanNumber Checks if given string is Roman numeral string. Returns true if string is empty.
// Supports numbers between 1 and 3999. Negative numbers and lowercase letters are unsupported and will throw an error.
func IsRomanNumber(romanNumeral string) bool {
	matched, _ := regexp.MatchString(`^M{0,3}(CM|CD|D?C{0,3})?(XC|XL|L?X{0,3})?(IX|IV|V?I{0,3})?$`, romanNumeral)

	return matched
}

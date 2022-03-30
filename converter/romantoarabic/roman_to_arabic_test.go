package romantoarabic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var successCases = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6,
	"VII": 7, "VIII": 8, "IX": 9, "X": 10, "XI": 11, "XII": 12,
	"XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17,
	"XVIII": 18, "XIX": 19, "XX": 20, "XXXI": 31, "XXXII": 32,
	"XXXIII": 33, "XXXIV": 34, "XXXV": 35, "XXXVI": 36, "XXXVII": 37,
	"XXXVIII": 38, "XXXIX": 39, "XL": 40, "XLI": 41, "XLII": 42,
	"XLIII": 43, "XLIV": 44, "XLV": 45, "XLVI": 46, "XLVII": 47,
	"XLVIII": 48, "XLIX": 49, "L": 50, "LXXXIX": 89, "XC": 90,
	"XCI": 91, "XCII": 92, "XCIII": 93, "XCIV": 94, "XCV": 95,
	"XCVI": 96, "XCVII": 97, "XCVIII": 98, "XCIX": 99, "C": 100,
	"CI": 101, "CII": 102, "CIII": 103, "CIV": 104, "CV": 105,
	"CVI": 106, "CVII": 107, "CVIII": 108, "CIX": 109, "CXLIX": 149,
	"CCCXLIX": 349, "CDLVI": 456, "D": 500, "DCIV": 604, "DCCLXXXIX": 789,
	"DCCCXLIX": 849, "CMIV": 904, "M": 1000, "MVII": 1007, "MLXVI": 1066,
	"MCCXXXIV": 1234, "MDCLXVI": 1666, "MDCCLXXVI": 1776, "MMXXI": 2021, "MMDCCCVI": 2806,
	"MMCMXCIX": 2999, "MMM": 3000, "MMMCMLXXIX": 3979, "MMMCMXCIX": 3999,
}

var errorCases = []string{"IVCMXCIX", "IIIIII", "CCCMMVIIVV", "56", "ii", "iV", "-I"}

func TestToInteger(t *testing.T) {
	for input, expected := range successCases {
		out, err := ToInteger(input)
		assert.Nil(t, err)
		assert.Equal(t, out, expected)
	}

	for _, input := range errorCases {
		output, err := ToInteger(input)
		assert.Equal(t, output, 0)
		assert.Error(t, err)
	}

	val, err := ToInteger("")
	assert.Equal(t, val, 0)
	assert.Nil(t, err)
}

func TestIsRomanNumber(t *testing.T) {
	for input, _ := range successCases {
		isRomanNumber := IsRomanNumber(input)
		assert.True(t, isRomanNumber)
	}

	for _, input := range errorCases {
		isRomanNumber := IsRomanNumber(input)
		assert.False(t, isRomanNumber)
	}

	isRomanNumber := IsRomanNumber("")
	assert.True(t, isRomanNumber)
}

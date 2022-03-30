package model

type ConvertRomanNumeral struct {
	Input string `binding:"required,max=20,is_roman_numeral" json:"input"`
}

package propetybased

import "strings"

type Dictionary struct {
	arabic int
	roman  string
}

var romanSymbwolsBase = []Dictionary{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range romanSymbwolsBase {
		for arabic >= numeral.arabic {
			result.WriteString(numeral.roman)
			arabic -= numeral.arabic
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	var result int

	for _, numeral := range romanSymbwolsBase {
		for strings.HasPrefix(roman, numeral.roman) {
			result += numeral.arabic
			roman = strings.TrimPrefix(roman, numeral.roman)
		}
	}

	return result
}

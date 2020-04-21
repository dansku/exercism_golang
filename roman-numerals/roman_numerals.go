// Package romannumerals provides all the tools for the romannumerals exercism task.
package romannumerals

import "errors"

// ToRomanNumeral gets an arabic number and return a roman number
func ToRomanNumeral(number int) (string, error) {

	// Test for errors and boudaries
	if number <= 0 || number > 3000 {
		return "", errors.New("Unsupported number")
	}

	var romanReturn = ""

	// Get thousands
	var thousand = number / 1000

	for i := 0; i < thousand; i++ {
		romanReturn += "M"
	}

	var hundreds int = (number % 1000) / 100
	var decimals int = (number % 100000 % 100) / 10
	var units int = (number % 1000 % 100 % 10)

	if hundreds > 0 {
		romanReturn += romanMath(hundreds, [4]string{"C", "CD", "D", "CM"})
	}

	if decimals > 0 {
		romanReturn += romanMath(decimals, [4]string{"X", "XL", "L", "XC"})
	}

	if units > 0 {
		romanReturn += romanMath(units, [4]string{"I", "IV", "V", "IX"})
	}

	return romanReturn, nil

}

func romanMath(number int, values [4]string) string {
	romanReturn := ""

	switch number {
	case 1, 2, 3:
		for i := 0; i < number; i++ {
			romanReturn += values[0] // I
		}
	case 4:
		romanReturn += values[1] // "IV"
	case 5, 6, 7, 8:
		romanReturn += values[2] //"V"
		for i := 0; i < number-5; i++ {
			romanReturn += values[0] //"I"
		}
	default:
		romanReturn += values[3] // "IX"
	}

	return romanReturn
}

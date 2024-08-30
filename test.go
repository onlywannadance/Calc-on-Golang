package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func main() {

	var num1, num2 string
	var sign string
	fmt.Printf("Ввод: \n")

	n, err := fmt.Scanln(&num1, &sign, &num2)
	if err != nil || n > 3 {
		panic("Ошибка! Формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	} else {
		if IsLetter(num1) && IsLetter(num2) {
			result := calculateForRoman(num1, num2, sign)
			fmt.Println(result)
		} else {
			result := calculateForArab(num1, num2, sign)
			fmt.Println(result)
		}
	}
}

func calculateForRoman(num1, num2, sign string) string {

	var result int

	number1 := convertFromRomanToArab(num1)
	number2 := convertFromRomanToArab(num2)

	if checkNumber(number1) && checkNumber(number2) {

		switch sign {
		case "+":
			result = number1 + number2
		case "-":
			result = number1 - number2
		case "*":
			result = number1 * number2
		case "/":
			if number2 != 0 {
				result = number1 / number2
			} else {
				panic("Ошибка! Введены числа больше 10")
			}
		default:
			panic("Ошибка! Введены некорректные данные")
		}
	} else {
		panic("Введены числа больше 10")
	}
	if result <= 0 {
		panic("Ошибка! Некорректный результат")
	}

	res := convertFromArabToRoman(result)

	return res
}

func calculateForArab(num1, num2, sign string) int {

	var result int

	number1, err := strconv.Atoi(num1)
	if err != nil {
		panic("Ошибка! Введены некорректные данные")
	}
	number2, err := strconv.Atoi(num2)
	if err != nil {
		panic("Ошибка! Введены некорректные данные")
	}

	if checkNumber(number1) && checkNumber(number2) {
		switch sign {
		case "+":
			result = number1 + number2
		case "-":
			result = number1 - number2
		case "*":
			result = number1 * number2
		case "/":
			if number1/number2 != 0 {
				result = number1 / number2
			} else {
				panic("Ошибка! Введены некорректные данные")
			}
		default:
			panic("Ошибка! Некорректный знак")
		}
	} else {
		panic("Введены числа больше 10")
	}
	return result
}

func convertFromArabToRoman(number int) string {
	conversions := []struct {
		arab  int
		roman string
	}{
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

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.arab {
			roman += conversion.roman
			number -= conversion.arab
		}
	}
	return roman
}

func convertFromRomanToArab(roman string) int {
	convertRoman := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var arabNum, tmpNum int
	for i := len(roman) - 1; i >= 0; i-- {
		romanDigit := roman[i]
		arabDigit := convertRoman[romanDigit]
		if arabDigit < tmpNum {
			arabNum -= arabDigit
		} else {
			arabNum += arabDigit
			tmpNum = arabDigit
		}
	}
	return arabNum
}

func IsLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func checkNumber(number int) bool {
	return number <= 10
}

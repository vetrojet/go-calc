package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romans = []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
var operators = []string{"+", "-", "/", "*"}

func inArray(arr []string, arg string) bool {
	for _, value := range arr {
		if value == arg {
			return true
		}
	}
	return false
}

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
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

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func isRoman(arr []string) (bool, string) {

	if inArray(romans, arr[0]) && inArray(romans, arr[2]) {
		return true, ""
	} else if !inArray(romans, arr[0]) && !inArray(romans, arr[2]) {
		return false, ""
	}

	return false, "используются одновременно разные системы счисления"
}

func getRoman(val string) (int, string) {

	for i, roman := range romans {
		if val == roman {
			return i + 1, ""
		}
	}

	return 0, "формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)"
}

func calculate(arr []string) string {
	var result, firstNum, secondNum int
	var operator string
	var getIntEr error

	if len(arr) != 3 || !inArray(operators, arr[1]) {
		return "формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)"
	}

	isRoman, err := isRoman(arr)

	operator = arr[1]

	if !isRoman {
		firstNum, getIntEr = strconv.Atoi(arr[0])
		secondNum, getIntEr = strconv.Atoi(arr[2])

		if getIntEr != nil {
			return "формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)"
		}
	} else {
		firstNum, err = getRoman(arr[0])
		secondNum, err = getRoman(arr[2])
	}

	if err != "" {
		return err
	}

	if operator == "+" {
		result = firstNum + secondNum
	} else if operator == "-" {
		result = firstNum - secondNum
	} else if operator == "/" {
		result = firstNum / secondNum
	} else if operator == "*" {
		result = firstNum * secondNum
	}

	if isRoman {
		if result > 0 {
			return integerToRoman(result)
		} else {
			return "в римской системе нет отрицательных чисел."
		}
	}
	return fmt.Sprintf("%v", result)

}

func main() {

	for {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {

			SplitText := strings.Split(scanner.Text(), " ")
			result := calculate(SplitText)
			fmt.Println(result)

		}

	}
}

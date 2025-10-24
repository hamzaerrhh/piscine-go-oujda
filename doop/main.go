package main

import (
	"os"
)

const (
	MaxInt32 = 2147483647
	MinInt32 = -2147483648
)

func main() {
	if len(os.Args) != 4 {
		return
	}

	operator := os.Args[2]

	if !IsNumber(os.Args[1]) || !IsNumber(os.Args[3]) {
		return
	}

	num1, err1 := ConvertNumber(os.Args[1])
	num2, err2 := ConvertNumber(os.Args[3])
	if err1 != "" || err2 != "" {
		return
	}

	res := 0
	err := ""
	switch operator {
	case "+":
		res = num1 + num2

		if res-num1 != num2 {
			return
		}
	case "-":
		res = num1 - num2

		if res+num2 != num1 {
			return
		}
	case "/":
		if num2 == 0 {
			err = "No division by 0\n"
		} else {
			res = num1 / num2
		}
	case "%":

		if num2 == 0 {
			err = "No modulo by 0\n"
		} else {
			res = num1 % num2
		}
	case "*":
		res = num1 * num2
		if num1 == 0 || num2 == 0 {
			res = 0
			break
		}
		if res/num1 == num2 {
			res = num1 * num2
		} else {
			return
		}
	default:
		return
	}
	resultString := convertToString(res)
	// print err or res
	Print(resultString, err)
}

func IsNumber(number string) bool {
	if len(number) == 0 {
		return false
	}

	start := 0
	if number[0] == '+' || number[0] == '-' {
		if len(number) == 1 {
			return false // just "+" or "-" is invalid
		}
		start = 1
	}

	for _, char := range number[start:] {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func ConvertNumber(str string) (int, string) {
	if len(str) == 0 {
		return 0, "err" // or error handling
	}

	res := 0
	sign := 1
	start := 0

	// Handle sign
	if str[0] == '-' {
		sign = -1
		start = 1
	} else if str[0] == '+' {
		start = 1
	}

	for i := start; i < len(str); i++ {
		char := str[i]
		if char < '0' || char > '9' {
			break
		}

		digit := int(char - '0')

		if sign == 1 && (res > (MaxInt32-digit)/10) {
			return 0, "error"
		}
		if sign == -1 && (res > (-(MinInt32 + digit))/10) { // careful with MinInt32 here
			return 0, "error"
		}

		res = res*10 + digit
	}

	return res * sign, ""
}

func convertToString(nbr int) string {
	if nbr == 0 {
		return "0\n"
	}

	res := ""
	isNegative := false

	if nbr < 0 {
		isNegative = true
		nbr = -nbr // make nbr positive for digit extraction
	}

	for nbr > 0 {
		res = string(rune(nbr%10+'0')) + res
		nbr /= 10
	}

	if isNegative {
		res = "-" + res
	}

	return res + "\n"
}

func Print(result string, err string) {
	if err != "" {
		os.Stdout.Write([]byte(err))
		return
	}
	os.Stdout.Write([]byte(result))
}

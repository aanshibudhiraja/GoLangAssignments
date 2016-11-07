package frequencyCalculator

import "fmt"

// PrintNumberfrequency method
func PrintNumberfrequency(number int) bool {
	sl := make([]int, 10, 10)

	var singleDigit int

	var isValidNumber = false

	for number != 0 {
		singleDigit = number % 10
		sl[singleDigit]++
		number = number / 10
		if !isValidNumber {
			isValidNumber = true
		}
	}

	if isValidNumber {
		for k, v := range sl {
			fmt.Println(k, " : ", v)
		}
	} else {
		fmt.Println("Invalid or out of range input")
	}
	return isValidNumber
}

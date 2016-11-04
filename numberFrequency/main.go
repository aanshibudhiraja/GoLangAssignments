package main

import (
	"fmt"
	"numberFrequency/frequencyCalculator"
)

func main() {
	var number int
	fmt.Scanf("%d", &number)
	frequencyCalculator.PrintNumberfrequency(number)
}

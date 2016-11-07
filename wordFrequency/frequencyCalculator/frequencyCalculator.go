package frequencyCalculator

import (
	"fmt"
	"strings"
)

// PrintRepeatedCount method to print something
func PrintRepeatedCount(entireString string) map[string]int {

	allWords := strings.Split(entireString, " ")

	wordsMap := make(map[string]int)

	for i := 0; i < len(allWords); i++ {
		wordsMap[strings.ToLower(allWords[i])]++
	}
	for k, v := range wordsMap {
		fmt.Println(k, ":", v)
	}
	return wordsMap
}

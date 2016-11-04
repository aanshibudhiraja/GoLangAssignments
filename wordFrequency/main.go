package main

import (
	"bufio"
	"os"
	"strings"
	"wordFrequency/frequencyCalculator"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	userText, _ := reader.ReadString('\n')
	removedNewLineText := strings.Split(userText, "\n")
	//Pass the string entered by the user, not \n (enter)
	frequencyCalculator.PrintRepeatedCount(removedNewLineText[0])
}

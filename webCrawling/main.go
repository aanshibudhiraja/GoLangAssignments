package main

import (
	"fmt"
	webCrawling "webCrawling/crawlingOperation"
)

func main() {

	pythonLang := webCrawling.Lang{
		Name:      "Python",
		URL:       "https://www.python.org/",
		Bytes:     0,
		TimeTaken: 0,
	}
	rubyLang := webCrawling.Lang{
		Name:      "ruby",
		URL:       "https://www.ruby-lang.org/en/",
		Bytes:     0,
		TimeTaken: 0,
	}
	goLang := webCrawling.Lang{
		Name:      "go",
		URL:       "https://golang.org/",
		Bytes:     0,
		TimeTaken: 0,
	}

	fmt.Println("\nCrawling data in Normal format:")

	// Print in normal format
	webCrawling.PrintCrawlingTimeInNormalFormat(pythonLang)
	webCrawling.PrintCrawlingTimeInNormalFormat(rubyLang)
	webCrawling.PrintCrawlingTimeInNormalFormat(goLang)

	fmt.Println("\n\nCrawling data in JSON format:")
	// Printin JSON format
	webCrawling.PrintCrawlingTimeInJSONFormat(pythonLang)
	webCrawling.PrintCrawlingTimeInJSONFormat(rubyLang)
	webCrawling.PrintCrawlingTimeInJSONFormat(goLang)
}

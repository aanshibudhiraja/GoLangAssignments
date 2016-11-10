package main

import (
	"fmt"
	webCrawling "webCrawling/crawlingOperation"
)

func main() {

	pythonLang := webCrawling.Lang{
		Name: "Python",
		URL:  "https://www.python.org/",
	}
	rubyLang := webCrawling.Lang{
		Name: "ruby",
		URL:  "https://www.ruby-lang.org/en/",
	}
	goLang := webCrawling.Lang{
		Name: "go",
		URL:  "https://golang.org/",
	}

	fmt.Println("\nCrawling data in Normal format:")

	// Print in normal format
	// webCrawling.PrintCrawlingTimeInNormalFormat(pythonLang)
	// webCrawling.PrintCrawlingTimeInNormalFormat(rubyLang)
	// webCrawling.PrintCrawlingTimeInNormalFormat(goLang)

	fmt.Println("\n\nCrawling data in JSON format:")
	// Printin JSON format
	webCrawling.PrintCrawlingTimeInJSONFormat(pythonLang)
	webCrawling.PrintCrawlingTimeInJSONFormat(rubyLang)
	webCrawling.PrintCrawlingTimeInJSONFormat(goLang)
}

package main

import (
	"fmt"
	"time"
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

	ch := make(chan webCrawling.Lang)

	start := time.Now()

	// Print in normal forma
	go webCrawling.PrintCrawlingTimeInNormalFormat(pythonLang, ch)
	pythonLang = <-ch
	go webCrawling.PrintCrawlingTimeInNormalFormat(rubyLang, ch)
	rubyLang = <-ch
	go webCrawling.PrintCrawlingTimeInNormalFormat(goLang, ch)
	goLang = <-ch

	totalTime := time.Since(start)

	fmt.Printf("\nTotal time taken : %s \nSome of time taken by all links: %fs\n", totalTime, (pythonLang.TimeTaken.Seconds() + rubyLang.TimeTaken.Seconds() + goLang.TimeTaken.Seconds()))

	fmt.Println("\n\nCrawling data in JSON format:")
	// Printin JSON format

	start = time.Now()

	go webCrawling.PrintCrawlingTimeInJSONFormat(pythonLang, ch)
	pythonLang = <-ch
	go webCrawling.PrintCrawlingTimeInJSONFormat(rubyLang, ch)
	rubyLang = <-ch
	go webCrawling.PrintCrawlingTimeInJSONFormat(goLang, ch)
	goLang = <-ch

	totalTime = time.Since(start)

	fmt.Printf("\nTotal time taken : %s\nSome of time taken by all links: %fs\n", totalTime, (pythonLang.TimeTaken.Seconds() + rubyLang.TimeTaken.Seconds() + goLang.TimeTaken.Seconds()))

}

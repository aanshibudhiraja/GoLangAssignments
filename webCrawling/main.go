package main

import (
	"fmt"
	"sync"
	"time"
	webCrawling "webCrawling/crawlingOperation"
)

func main() {

	var wg sync.WaitGroup

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
	wg.Add(3)
	start := time.Now()
	go webCrawling.PrintCrawlingTimeInNormalFormat(&pythonLang, &wg)
	go webCrawling.PrintCrawlingTimeInNormalFormat(&rubyLang, &wg)
	go webCrawling.PrintCrawlingTimeInNormalFormat(&goLang, &wg)

	wg.Wait()

	totalTimeTaken := time.Since(start)

	fmt.Printf("\nTotal time taken : %s\nSome of time taken by all links: %f\n", totalTimeTaken, (pythonLang.TimeTaken.Seconds() + rubyLang.TimeTaken.Seconds() + goLang.TimeTaken.Seconds()))

	fmt.Println("\n\nCrawling data in JSON format:")
	// Printin JSON format
	wg.Add(3)
	start = time.Now()

	webCrawling.PrintCrawlingTimeInJSONFormat(&pythonLang, &wg)
	webCrawling.PrintCrawlingTimeInJSONFormat(&rubyLang, &wg)
	webCrawling.PrintCrawlingTimeInJSONFormat(&goLang, &wg)

	wg.Wait()
	totalTimeTaken = time.Since(start)

	fmt.Printf("\nTotal time taken : %s\nSome of time taken by all links: %f\n", totalTimeTaken, (pythonLang.TimeTaken.Seconds() + rubyLang.TimeTaken.Seconds() + goLang.TimeTaken.Seconds()))

}

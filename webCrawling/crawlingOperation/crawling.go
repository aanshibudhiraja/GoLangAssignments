package webCrawling

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

// Lang type
type Lang struct {
	Name      string
	URL       string
	Bytes     int64
	TimeTaken time.Duration
}

// PrintCrawlingTimeInNormalFormat method
func PrintCrawlingTimeInNormalFormat(emptyLang *Lang, wg *sync.WaitGroup) {
	calculateCrawlinngTime(func(lang Lang) {
		fmt.Printf("\n%v", lang)
	}, emptyLang)
	wg.Done()
}

// PrintCrawlingTimeInJSONFormat method
func PrintCrawlingTimeInJSONFormat(emptyLang *Lang, wg *sync.WaitGroup) {
	calculateCrawlinngTime(func(lang Lang) {
		b, _ := json.Marshal(lang)
		// Convert bytes to string.
		s := string(b)
		fmt.Println(s)
	}, emptyLang)
	wg.Done()
}

// calculateCrawlinngTime method
func calculateCrawlinngTime(printFunc func(Lang), lang *Lang) {
	start := time.Now()
	data, _ := http.Get(lang.URL)
	lang.TimeTaken = time.Since(start)
	lang.Bytes, _ = io.Copy(ioutil.Discard, data.Body)
	printFunc(*lang)
}

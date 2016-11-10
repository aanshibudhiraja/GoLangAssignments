package webCrawling

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Lang type
type Lang struct {
	Name      string
	URL       string
	Bytes     int64
	TimeTaken string
}

// PrintCrawlingTimeInNormalFormat method
func PrintCrawlingTimeInNormalFormat(lang Lang) {
	calculateCrawlinngTime(func(lang Lang) {
		fmt.Printf("\n%v", lang)
	}, lang)
}

// PrintCrawlingTimeInJSONFormat method
func PrintCrawlingTimeInJSONFormat(lang Lang) {
	calculateCrawlinngTime(func(lang Lang) {
		b, _ := json.Marshal(lang)
		// Convert bytes to string.
		s := string(b)
		fmt.Println(s)
	}, lang)
}

// calculateCrawlinngTime method
func calculateCrawlinngTime(printFunc func(Lang), lang Lang) {
	start := time.Now()
	data, _ := http.Get(lang.URL)
	lang.TimeTaken = time.Since(start).String()
	lang.Bytes, _ = io.Copy(ioutil.Discard, data.Body)
	printFunc(lang)
}

package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"vz",
		"t",
		"tmus",
		"s",
	}
	// in  this situaion we know how many go routines will be running
	// that is number of symbol in slice
	numComplete := 0
	for _, sybmol := range stockSymbols {
		go func(sybmol string) {

			resp, _ := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=" + sybmol)
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)
			quote := new(QuoteResponse)
			xml.Unmarshal(body, &quote)
			fmt.Printf("%s : %.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(sybmol)
	}
	for numComplete < len(stockSymbols) {
		time.Sleep(10 * time.Millisecond)
	}
	elapsed := time.Since(start)
	fmt.Printf("execution time : %s", elapsed)
	// go func() {
	// 	println("hello")
	// }()
	// go func() {
	// 	println("go")
	// }()
}

type QuoteResponse struct {
	Status        string
	Name          string
	LastPrice     float32
	Change        float32
	ChangePercent float32
	TimeStamp     string
	MSDate        float32
	Volume        int
}

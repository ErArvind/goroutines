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
	resp, _ := http.Get("http://dev.markitondemand.com/MODApis/Api/v2/Quote?symbol=googl")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	quote := new(QuoteResponse)
	xml.Unmarshal(body, &quote)
	fmt.Printf("%s : %.2f", quote.Name, quote.LastPrice)
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

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

func getFlags() (string, int) {
	url := flag.String("URL", "https://httpbin.org/delay/3", "URL to open")
	iterations := flag.Int("iterations", 3, "Number of attempts")
	flag.Parse()
	return *url, *iterations
}

func main() {
	url, iterations := getFlags()
	for i := 0; i < iterations; i++ {
		timeout := math.Pow(2, float64(i))
		client := &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		}
		fmt.Printf("This is iteration number %d. The timeout is set to %d", i+1, int(timeout))
		res, err := client.Get(url)
		if err != nil {
			fmt.Println(" ... request failed")
			fmt.Println(err)
		} else {
			data, err := ioutil.ReadAll(res.Body)
			res.Body.Close()
			if err != nil {
				fmt.Println(" ... request succeeded")
				fmt.Println(err)
			} else {
				fmt.Println(" ... request succeeded")
				fmt.Println(string(data))
				break
			}
		}
	}
}
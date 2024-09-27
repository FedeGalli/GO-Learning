package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var maxPrice float32 = 10.0
var wg = sync.WaitGroup{}

func main() {
	t0 := time.Now()
	c := make(chan float32, 3)
	availableSites := []string{"amazon.com", "lidl.com", "conad.it"}
	for _, site := range availableSites {
		wg.Add(1)
		go checkGPUPrices(site, c)
	}
	wg.Wait()
	fmt.Printf("Found all after %v", time.Since(t0))
}

func checkGPUPrices(site string, channel chan float32) {
	for {
		waitTime := rand.ExpFloat64() * 2000
		time.Sleep(time.Millisecond * time.Duration(waitTime)) //checking the price every x seconds
		currentPrice := rand.Float32() * 20

		if currentPrice <= maxPrice {
			channel <- currentPrice
			sendNotification(site, channel)
			wg.Done()
			break
		}
	}
}

func sendNotification(site string, channel chan float32) {
	fmt.Println("Found a new GPU at "+site+" price: %v", <-channel)
}

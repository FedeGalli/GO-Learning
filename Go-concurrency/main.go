package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var db = []int{1, 2, 5, 10, 12}
var total = 0
var channel = make(chan int)
var wg = sync.WaitGroup{}

func main() {
	t0 := time.Now()
	wg.Add(1)
	go readDB(db)
	go sum()
	wg.Wait()
	fmt.Printf("Operation took: %v\n", time.Since(t0))
	fmt.Println(total)
}

func readDB(nums []int) {
	for _, i := range nums {
		time.Sleep(time.Duration(rand.Float64()*1000) * time.Millisecond)
		channel <- i
	}
	close(channel)
}

func sum() {
	for val := range channel {
		total += val
	}
	wg.Done()
}

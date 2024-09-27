package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var m = sync.Mutex{}
var dbData [4]string = [4]string{"1", "2", "3", "4"}
var res []string = []string{}

func main() {
	t0 := time.Now()

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go readDb(i)
	}
	wg.Wait()
	fmt.Printf("Operation took: %v", time.Since(t0))
	fmt.Println(res)
}

func readDb(i int) {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("Result is: " + dbData[i])
	m.Lock()
	res = append(res, dbData[i])
	m.Unlock()
	wg.Done()
}

func save() {

}

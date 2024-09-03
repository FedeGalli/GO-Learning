package main

import "fmt"

func main() {
	myString := "Hello worldè+"
	fmt.Println(myString)

	for i, char := range myString {
		fmt.Printf("index: %v char: %v\n", i, char)
	}

	myRune := []rune("Hèll+ worldè")

	for i, char := range myRune {
		fmt.Printf("index: %v char: %v\n", i, char)
	}
}

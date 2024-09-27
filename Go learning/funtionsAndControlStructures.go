package main

import (
	"errors"
	"fmt"
)

func main() {
	num := 10
	den := 2
	division, rest, err := intDivision(num, den)

	if err == nil {

	} else {
		fmt.Printf(err.Error())
	}

	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case rest == 0:
		fmt.Printf("The rest is %v, perfect division", rest)
	default:
		fmt.Printf("Division: %v \nRest: %v", division, rest)
	}

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(num int, den int) (int, int, error) {
	var err error
	if den == 0 {
		err = errors.New("Cannot divide by 0")
		return -1, -1, err
	}
	return (num / den), num % den, err
}

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//for each datatype the default value is:
	//Numeric variables = 0
	//String variables = ""
	//Boolean variables = false

	var intNum int16 = 10
	fmt.Println(intNum)

	var floatNum float64 = 10.12354123
	fmt.Println(floatNum)

	var result float64 = float64(intNum) + floatNum
	fmt.Println(result)

	var s string = "Hello!\n" + " Worldδ"
	fmt.Println(len(s)) //character outside the vanilla ascii character set are counted double, \n is counted too

	//to avoid this
	fmt.Println(utf8.RuneCountInString(s))

	var myBool bool = false
	fmt.Println(myBool)

	fastBool := false //fast initialization with auto-infering

	fastBool, myBool = true, true

	fmt.Println(myBool, fastBool)

	const myConst int8 = 10
}

package main

import "fmt"

func main() {
	//arrays()
	//slices()
	//maps()
	loops()

}

func arrays() {
	/*
		Arays are:
			*FIXED SIZE
			*SAME TYPE
			*INDEXABLE
			*CONTIGUOUS IN MEMORY
	*/
	var intArr [3]int32 = [3]int32{1, 2, 3} //or intArr := [3]int32{1, 2, 3}
	intArr[2] = 1
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	//to get the memory location use the & char
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])
}

func slices() {

	//Slices
	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	//append return a new slice with the 7 appended
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)

	//the actual length of the slice
	fmt.Println(len(intSlice))
	//the number of potential values that can be stored (it dinamically adjust based on the len)
	fmt.Println(cap(intSlice))

	intSlice2 := []int32{10, 11, 12}

	fmt.Println(append(intSlice, intSlice2...))

	//to initialize a slice we can use make(type, len, cap) function as well
	intSlice3 := make([]int32, 3, 8)
	fmt.Println(intSlice3)
}

func maps() {
	//set of key value pair
	myMap := map[string]uint8{"Adam": 10, "Carlos": 12}

	fmt.Println(myMap["Carlos"])

	//carefull because maps are always to return a value, even if the key does not exist, in that case the default value is returned

}

func loops() {
	myMap := map[string]uint8{"Adam": 10, "Carlos": 12}
	for name, age := range myMap {
		fmt.Println(name, age)
	}

	//while loop
	i := 0
	for i < 10 {
		fmt.Println("ciao")
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println("ciao")
	}
}

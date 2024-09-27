package main

import "fmt"

func square(nums *[5]float64) [5]float64 {
	for i, num := range nums {
		nums[i] = num * num
	}
	return *nums
}

func main() {
	var p *int32 = new(int32)
	var i int32

	p = &i

	i = 100

	*p = 200

	fmt.Println(p, &i)

	nums := [5]float64{1.1, 2.2, 3.3, 4.4, 5.5}
	q := square(&nums)
	fmt.Println(q[0])

}

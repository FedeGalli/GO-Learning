package arrayandslices

func Add(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func AddAll(n1, n2 []int) []int {
	s1 := 0
	s2 := 0
	for _, num := range n1 {
		s1 += num
	}
	for _, num := range n2 {
		s2 += num
	}
	return []int{s1, s2}
}

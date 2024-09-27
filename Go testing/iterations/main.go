package iterations

func Repeat(char string, times int) string {
	s := ""
	for i := 0; i < times; i++ {
		s += char
	}

	return s
}

package my

func Reverse(s *string) string {
	res := ""
	for _, char := range *s {
		res += string(char) + res
	}

	return res
}

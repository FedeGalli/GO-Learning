package usage

import (
	"fmt"

	my "github.com/fedegalli/mystring"
)

func Main() {
	s := "ciao"
	fmt.Println(my.Reverse(&s))
}

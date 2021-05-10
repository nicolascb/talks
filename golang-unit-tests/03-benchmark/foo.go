package foo

import "fmt"

func A(s string) string {
	for i := 0; i <= 5; i++ {
		s = fmt.Sprintf("%s-%d", s, i)
	}

	return s
}

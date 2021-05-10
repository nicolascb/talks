package foo

func Foo(s string) string {
	b := b(s)
	c := c(s)
	return b + "_" + c
}

func b(s string) string {
	return s + "_b"
}

func c(s string) string {
	return s + "_c"
}

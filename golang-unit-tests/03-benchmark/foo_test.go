package foo_test

import (
	"foo"
	"testing"
)

func TestA(t *testing.T) {
	arg := "hello"
	expected := "hello-0-1-2-3-4-5"

	if res := foo.A(arg); res != expected {
		t.Errorf("received: %s, expected: %s", res, expected)
	}
}

func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		foo.A("hello")
	}
}

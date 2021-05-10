package foo_test

import (
	"foo"
	"reflect"
	"testing"
)

func TestFoo(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args func(t *testing.T) args

		want1 string
	}{
		{
			name: "FooString",
			args: func(*testing.T) args {
				return args{
					s: "hello",
				}
			},
			want1: "hello_b_hello_c",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1 := foo.Foo(tArgs.s)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Foo got1 = %v, want1: %v", got1, tt.want1)
			}
		})
	}
}

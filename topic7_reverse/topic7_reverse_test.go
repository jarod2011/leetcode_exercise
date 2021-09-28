package topic7_reverse

import (
	"fmt"
	"testing"
)

func Test_reverseByConvertString(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		args args
		want int
	}{
		{args: args{x: 123}, want: 321},
		{args: args{x: 123}, want: 321},
		{args: args{x: 1113210}, want: 123111},
		{args: args{x: -123}, want: -321},
		{args: args{x: 123456789120}, want: 0},
		{args: args{x: -1234567893}, want: 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("test for arg %#v want %#v", tt.args, tt.want), func(t *testing.T) {
			if got := reverseByConvertString(tt.args.x); got != tt.want {
				t.Errorf("reverseUseString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_reverseByConvertString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseByConvertString(123)
	}
}

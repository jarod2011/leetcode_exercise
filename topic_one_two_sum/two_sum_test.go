package topic_one_two_sum

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		args args
		want []int
	}{
		{args: args{nums: []int{2,7,11,15}, target: 9}, want: []int{0, 1}},
		{args: args{nums: []int{3,2,4}, target: 6}, want: []int{1, 2}},
		{args: args{nums: []int{3,3}, target: 6}, want: []int{0, 1}},
		{args: args{nums: []int{2,7,11,15}, target: 16}, want: []int{}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("test args %#v want %#v", tt.args, tt.want), func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_twoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		twoSum([]int{2,7,11,15}, 9)
	}
}
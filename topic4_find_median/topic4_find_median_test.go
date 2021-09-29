package topic4_find_median

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		args args
		want float64
	}{
		{args: args{nums1: []int{1, 3}, nums2: []int{2}}, want: 2},
		{args: args{nums1: []int{1, 2}, nums2: []int{3, 4}}, want: 2.5},
		{args: args{nums1: []int{0, 0}, nums2: []int{0, 0}}, want: 0},
		{args: args{nums1: []int{}, nums2: []int{1}}, want: 1},
		{args: args{nums1: []int{2}, nums2: []int{}}, want: 2},
		{args: args{nums1: []int{1, 3, 4, 9}, nums2: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}}, want: 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("test for args %#v want %#v", tt.args, tt.want), func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findMedianSortedArrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		findMedianSortedArrays([]int{1, 3, 4, 9}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	}
}

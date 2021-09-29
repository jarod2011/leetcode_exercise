package topic4_find_median

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	l := len(nums1) + len(nums2)
	if l%2 == 1 {
		return float64(findLessK(nums1, nums2, l/2+1))
	}
	return float64(findLessK(nums1, nums2, l/2)+findLessK(nums1, nums2, l/2+1)) / 2
}

func findLessK(nums1 []int, nums2 []int, k int) int {
	var offset1, offset2, index1, index2 int
	for {
		if len(nums1) == offset1 {
			return nums2[offset2+k-1]
		}
		if len(nums2) == offset2 {
			return nums1[offset1+k-1]
		}
		if k == 1 {
			return minInt(nums1[offset1], nums2[offset2])
		}
		index1 = minInt(offset1+k/2, len(nums1)) - 1
		index2 = minInt(offset2+k/2, len(nums2)) - 1
		if nums1[index1] < nums2[index2] {
			k -= index1 - offset1 + 1
			offset1 = index1 + 1
		} else {
			k -= index2 - offset2 + 1
			offset2 = index2 + 1
		}
	}
}

func minInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

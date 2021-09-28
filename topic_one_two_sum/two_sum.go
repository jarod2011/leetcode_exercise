package topic_one_two_sum

func twoSum(nums []int, target int) []int {
	var cache = make(map[int]int)
	for k, v := range nums {
		if idx, ok := cache[v]; ok {
			return []int{idx, k}
		}
		cache[target - v] = k
	}
	return []int{}
}
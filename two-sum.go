package main

func twoSum(nums []int, target int) []int {
	var cache = map[int]int{}
	var result []int
	for i, val := range nums {
		var delta = target - val
		if j, ex := cache[delta]; !ex {
			cache[val] = i
		} else {
			result = []int{i, j}
			break
		}
	}
	return result
}

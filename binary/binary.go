package binary

func GetTestData() ([]int, int) {
	return []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13}, 11
}

func getTestData2() ([]int, int) {
	return []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 0, 1, 2}, 1
}

func Search(nums []int, target int) int {
	var start = 0
	var end = len(nums)
	for start <= end {
		var mid = start + (end-start)/2
		if target == nums[mid] {
			return mid
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}

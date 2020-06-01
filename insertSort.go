package goAlgorithm

func InsertSort(nums []int) []int {
	n := len(nums)
	for i := 1; i < n; i++ {
		temp := nums[i]
		j := i
		for j > 0 && nums[j - 1] > temp {
			nums[j] = nums[j - 1]
			j--
		}
		nums[j] = temp
	}
	return nums
}

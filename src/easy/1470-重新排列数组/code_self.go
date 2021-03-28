package _470_重新排列数组

func shuffle(nums []int, n int) []int {
	var result = make([]int, len(nums))
	left, right := nums[:n], nums[n:]
	for i := 0; i < n; i++ {
		result[i*2] = left[i]
		result[i*2+1] = right[i]
	}
	return result
}

package sol

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func rob(nums []int) int {
	total := len(nums)

	var rob_with_range = func(start, end int) int {
		prevTwo := 0
		prevOne := nums[end]
		totalRob := 0
		for robIdx := end - 1; robIdx >= start; robIdx-- {
			totalRob = max(nums[robIdx]+prevTwo, prevOne)
			prevTwo = prevOne
			prevOne = totalRob
		}
		return max(prevOne, prevTwo)
	}
	if total == 1 {
		return nums[0]
	}

	return max(rob_with_range(0, total-2), rob_with_range(1, total-1))
}

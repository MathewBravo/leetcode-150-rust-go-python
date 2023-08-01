package pkg


func RemoveDuplicates2(nums []int) int {
	var left, right int
	for right < len(nums) {
		count := 1
		for right+1 < len(nums) && nums[right] == nums[right+1] {
			right++
			count++
		}
		min := 2
		if count < min {
			min = count
		}

		i := 0
		for i < min {
			nums[left] = nums[right]
			left++
			i++
		}
		right++
	}
	return left
}


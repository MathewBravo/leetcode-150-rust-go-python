package pkg

func rotate(nums []int, k int) {
	k = k % len(nums)

	reverse_in_place := func(nums []int, left, right int) {
		for left < right {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}

	reverse_in_place(nums, 0, len(nums)-1)
	reverse_in_place(nums, 0, k-1)
	reverse_in_place(nums, k, len(nums)-1)
}

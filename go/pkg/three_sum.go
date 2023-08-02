package pkg

func threeSum(nums []int) [][]int {
	var results [][]int
	var quicksort func([]int) []int
	quicksort = func(nums []int) []int {
		if len(nums) <= 1 {
			return nums
		}

		pivot := nums[0]
		var lesser []int
		var greater []int

		for _, num := range nums[1:] {
			if num <= pivot {
				lesser = append(lesser, num)
			} else {
				greater = append(greater, num)
			}
		}

		lesser = quicksort(lesser)
		greater = quicksort(greater)
		return append(append(lesser, pivot), greater...)
	}
	nums = quicksort(nums)

	for i, num := range nums {
		if i > 0 && nums[i-1] == num {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sumTotal := num + nums[left] + nums[right]
			if sumTotal > 0 {
				right--
			} else if sumTotal < 0 {
				left++
			} else {
				results = append(results, []int{num, nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left-1] == nums[left] {
					left++
				}
				for right > left && nums[right+1] == nums[right] {
					right--
				}
			}
		}
	}
	return results
}

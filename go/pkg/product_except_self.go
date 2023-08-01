package pkg

func ProductExceptSelf(nums []int) []int {
	len := len(nums)
	prefix := make([]int, len)
	suffix := make([]int, len)
	for i := range nums {
		prefix[i] = 1
		suffix[i] = 1
	}

	for i := 1; i < len; i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}
	for i := len - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] * nums[i+1]
	}

	result := make([]int, len)
	for i := range result {
		result[i] = prefix[i] * suffix[i]
	}

	return result
}

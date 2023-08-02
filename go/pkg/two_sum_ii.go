package pkg

func twoSumII(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left <= right {
		if numbers[right]+numbers[left] > target {
			right--
		} else if numbers[right]+numbers[left] < target {
			left++
		} else {
			return []int{left + 1, right + 1}
		}
	}
	return []int{}
}

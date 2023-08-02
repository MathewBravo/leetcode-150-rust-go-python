package pkg

func trap(height []int) int {
	left, right := 0, len(height)-1
	max_left, max_right, sum := height[left], height[right], 0
	for left < right {
		if max_left < max_right {
			left++
			water := max_left - height[left]
			if water > 0 {
				sum += water
			}
			max_left = max(max_left, height[left])
		} else {
			right--
			water := max_right - height[right]
			if water > 0 {
				sum += water
			}
			max_right = max(max_right, height[right])
		}
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

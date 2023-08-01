package pkg

func MaxArea(height []int) int {
	left, right, max_area := 0, len(height)-1, 0
	for left < right {
		current_area := height[left] * (right - left)
		if current_area > height[right]*(right-left) {
			current_area = height[right] * (right - left)
		}
		if max_area < current_area {
			max_area = current_area
		}

		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}
	return max_area
}

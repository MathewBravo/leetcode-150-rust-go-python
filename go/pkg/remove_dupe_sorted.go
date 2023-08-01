package pkg

func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return 1
	}
	u_id := 0
	for _, num := range nums {
		if num != nums[u_id] {
			u_id += 1
			nums[u_id] = num
		}
	}
	return u_id + 1
}

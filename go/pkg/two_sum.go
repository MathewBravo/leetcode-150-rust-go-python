package pkg

func twoSum(numbers []int, target int) []int {
	seen_num := make(map[int]int)
	for i, num := range numbers {
		seen_num[num] = i
	}

	for _, num := range seen_num {
		_, ok := seen_num[target-num]
		if ok {
			return []int{seen_num[num] + 1, seen_num[target-num] + 1}
		}
	}

	return []int{-1, -1}
}

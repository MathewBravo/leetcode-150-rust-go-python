
package pkg

func romanToInt(s string) int {
	var sum int
	values := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for i, ch := range s {
		cur_val := values[ch]
		var next_val int
		if i+1 < len(s) {
			next_val = values[rune(s[i+1])]
		}
		if cur_val < next_val {
			sum -= cur_val
		} else {
			sum += cur_val
		}
	}
	return sum
}

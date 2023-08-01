package pkg

import "testing"

func TestMaxProfit(t *testing.T) {
	test := MaxProfit([]int{7, 1, 5, 3, 6, 4})
	result := 5
	if result != test {
		t.Errorf("maxProfit([7,1,5,3,6,4] expected %d, returned %d)", result, test)
	}
}

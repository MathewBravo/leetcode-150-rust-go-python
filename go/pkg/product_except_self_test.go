package pkg

import "testing"

func TestProductExceptSelf(t *testing.T) {
	test := ProductExceptSelf([]int{1, 2, 3, 4})
	expect := []int{24, 12, 8, 6}
	for i, num := range test {
		if num != expect[i] {
			t.Errorf("Product Except Self Failed. Expected %v, Got %v", expect, test)
		}
	}
}

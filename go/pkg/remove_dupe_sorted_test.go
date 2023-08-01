package pkg

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	test_var := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	test := RemoveDuplicates(test_var)
	expect := []int{0, 1, 2, 3, 4}
	for i, num := range expect {
		if num != test_var[i] {
			t.Errorf("Remove Duplicates 1 failed, Expected: %v, Got: %v, with %d dupes", expect, test_var[:len(expect)], test)
		}
	}

}

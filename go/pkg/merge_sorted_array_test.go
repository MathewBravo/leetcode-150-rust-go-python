package pkg

import "testing"

func TestMerge(t *testing.T) {
	merge_arr := []int{1, 2, 3, 0, 0, 0}
	Merge(merge_arr, 3, []int{2, 5, 6}, 3)
	for i, num := range []int{1, 2, 2, 3, 5, 6} {
		if merge_arr[i] != num {
			t.Errorf("Merge failed. Expected %v, Got %v", []int{1, 2, 2, 3, 5, 6}, merge_arr)
		}
	}

}

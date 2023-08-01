package pkg

import "testing"

func TestStrStr(t *testing.T) {
	test := StrStr("sadbutsad", "sad")
	expect := 0
	if expect != test {
		t.Errorf("Find index first occ in string failed, expected %d got %d", expect, test)
	}
}

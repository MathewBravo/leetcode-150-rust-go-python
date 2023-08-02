package pkg

import "testing"

func TestRomanToInt(t *testing.T) {
	result := romanToInt("MCMXCIV")
	if result != 1994 {
		t.Errorf("Roman To Int: Expected 1994, Got %d", result)
	}
}

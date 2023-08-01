package pkg

import (
	"strings"
)

func LengthOfLastWord(s string) int {
	splits := strings.Split(s, " ")
	return len(splits[len(splits)-1])
}

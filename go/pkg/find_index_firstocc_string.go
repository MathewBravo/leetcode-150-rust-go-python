package pkg

func StrStr(haystack string, needle string) int {
	if haystack == needle {
		return 0
	}
	n_len := len(needle)
	for i, _ := range haystack {
		if haystack[i:i+n_len] == needle {
			return i
		}
	}

	return -1
}

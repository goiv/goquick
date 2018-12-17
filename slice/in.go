package slice

import "strings"

func InInt(needle int, haystack []int) bool {
	var exist bool
	for _, current := range haystack {
		if needle == current {
			exist = true
			break
		}
	}

	return exist
}

func InInt32(needle int32, haystack []int32) bool {
	var exist bool
	for _, current := range haystack {
		if needle == current {
			exist = true
			break
		}
	}

	return exist
}

func InInt64(needle int64, haystack []int64) bool {
	var exist bool
	for _, current := range haystack {
		if needle == current {
			exist = true
			break
		}
	}

	return exist
}

func InStr(needle string, haystack []string) bool {
	var exist bool
	needle = strings.ToLower(needle)
	for _, current := range haystack {
		if needle == strings.ToLower(current) {
			exist = true
			break
		}
	}

	return exist
}
package slice

import "strings"

func AppendNotRepeatedInt(slice []int, d int) []int {
	for _, current := range slice {
		if current == d {
			return slice
		}
	}

	return append(slice, d)
}

func AppendNotRepeatedInt32(slice []int32, d int32) []int32 {
	for _, current := range slice {
		if current == d {
			return slice
		}
	}

	return append(slice, d)
}

func AppendNotRepeatedInt64(slice []int64, d int64) []int64 {
	for _, current := range slice {
		if current == d {
			return slice
		}
	}

	return append(slice, d)
}

func AppendNotRepeatedStr(slice []string, d string) []string {
	d = strings.ToLower(d)
	for _, current := range slice {
		if strings.ToLower(current) == d {
			return slice
		}
	}

	return append(slice, d)
}

package slice

import "testing"

func TestInInt(t *testing.T) {
	tests := []struct{
		needle int
		haystack []int
		expected bool
	}{
		{1, []int{1, 2, 3, 4, 5}, true},
		{3, []int{1, 2, 3}, true},
		{5, []int{1, 2, 3, 4}, false},
	}

	for _, test := range tests {
		if got := InInt(test.needle, test.haystack); got != test.expected {
			t.Errorf("In(%d, %v) = %t, expected: %t", test.needle, test.haystack, got, test.expected)
		}
	}
}

func TestInInt32(t *testing.T) {
	tests := []struct{
		needle int32
		haystack []int32
		expected bool
	}{
		{int32(1), []int32{1, 2, 3, 4, 5}, true},
		{int32(3), []int32{1, 2, 3}, true},
		{int32(5), []int32{1, 2, 3, 4}, false},
	}

	for _, test := range tests {
		if got := InInt32(test.needle, test.haystack); got != test.expected {
			t.Errorf("In(%d, %v) = %t, expected: %t", test.needle, test.haystack, got, test.expected)
		}
	}
}

func TestInInt64(t *testing.T) {
	tests := []struct{
		needle int64
		haystack []int64
		expected bool
	}{
		{int64(1), []int64{1, 2, 3, 4, 5}, true},
		{int64(3), []int64{1, 2, 3}, true},
		{int64(5), []int64{1, 2, 3, 4}, false},
	}

	for _, test := range tests {
		if got := InInt64(test.needle, test.haystack); got != test.expected {
			t.Errorf("In(%d, %v) = %t, expected: %t", test.needle, test.haystack, got, test.expected)
		}
	}
}

func TestInStr(t *testing.T) {
	tests := []struct{
		needle string
		haystack []string
		expected bool
	}{
		{string("D"), []string{"E", "D", "D", "Y"}, true},
		{string("G"), []string{"G", "O"}, true},
		{string("k"), []string{"quick"}, false},
		{string("eddy"), []string{"eddycjy"}, false},
	}

	for _, test := range tests {
		if got := InStr(test.needle, test.haystack); got != test.expected {
			t.Errorf("In(%s, %v) = %t, expected: %t", test.needle, test.haystack, got, test.expected)
		}
	}
}
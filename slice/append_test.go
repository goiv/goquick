package slice

import "testing"

func TestAppendNotRepeatedInt(t *testing.T) {
	tests := []struct{
		d int
		slice []int
		expected []int
	}{
		{1, []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{3, []int{1, 2, 3}, []int{1, 2, 3}},
		{5, []int{1, 2, 3, 4}, []int{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		if got := AppendNotRepeatedInt(test.slice, test.d); len(test.expected) != len(got) {
			t.Errorf("In(%d, %v) = %v, expected: %v", test.d, test.slice, got, test.expected)
		}
	}
}

func TestAppendNotRepeatedInt32(t *testing.T) {
	tests := []struct{
		d int32
		slice []int32
		expected []int32
	}{
		{1, []int32{1, 2, 3, 4, 5}, []int32{1, 2, 3, 4, 5}},
		{3, []int32{1, 2, 3}, []int32{1, 2, 3}},
		{5, []int32{1, 2, 3, 4}, []int32{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		if got := AppendNotRepeatedInt32(test.slice, test.d); len(test.expected) != len(got) {
			t.Errorf("In(%d, %v) = %v, expected: %v", test.d, test.slice, got, test.expected)
		}
	}
}

func TestAppendNotRepeatedInt64(t *testing.T) {
	tests := []struct{
		d int64
		slice []int64
		expected []int64
	}{
		{1, []int64{1, 2, 3, 4, 5}, []int64{1, 2, 3, 4, 5}},
		{3, []int64{1, 2, 3}, []int64{1, 2, 3}},
		{5, []int64{1, 2, 3, 4}, []int64{1, 2, 3, 4, 5}},
	}

	for _, test := range tests {
		if got := AppendNotRepeatedInt64(test.slice, test.d); len(test.expected) != len(got) {
			t.Errorf("In(%d, %v) = %v, expected: %v", test.d, test.slice, got, test.expected)
		}
	}
}

func TestAppendNotRepeatedStr(t *testing.T) {
	tests := []struct{
		d string
		slice []string
		expected []string
	}{
		{"1", []string{"1", "2", "3", "4", "5"}, []string{"1", "2", "3", "4", "5"}},
		{"3", []string{"1", "2", "3"}, []string{"1", "2", "3"}},
		{"5", []string{"1", "2", "3", "4"}, []string{"1", "2", "3", "4", "5"}},
	}

	for _, test := range tests {
		if got := AppendNotRepeatedStr(test.slice, test.d); len(test.expected) != len(got) {
			t.Errorf("In(%s, %v) = %v, expected: %v", test.d, test.slice, got, test.expected)
		}
	}
}
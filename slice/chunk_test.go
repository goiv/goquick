package slice

import (
	"testing"
)

func TestChunkInt(t *testing.T) {
	tests := []struct{
		slice []int
		size int
		expected [][]int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, [][]int{{1, 2}, {3, 4}, {5}}},
		{[]int{1, 2, 3}, 5, [][]int{{1, 2, 3}}},
		{[]int{1, 2}, 1, [][]int{{1}, {2}}},
		{[]int{1, 2}, 0, [][]int{{1, 2}}},
	}

	for _, test := range tests {
		if got := ChunkInt(test.slice, test.size); len(test.expected) != len(got) {
			t.Errorf("In(%v, %d) = %v, expected: %v", test.slice, test.size, got, test.expected)
		}
	}
}

func TestChunkInt32(t *testing.T) {
	tests := []struct{
		slice []int32
		size int
		expected [][]int32
	}{
		{[]int32{1, 2, 3, 4, 5}, 2, [][]int32{{1, 2}, {3, 4}, {5}}},
		{[]int32{1, 2, 3}, 5, [][]int32{{1, 2, 3}}},
		{[]int32{1, 2}, 1, [][]int32{{1}, {2}}},
		{[]int32{1, 2}, 0, [][]int32{{1, 2}}},
	}

	for _, test := range tests {
		if got := ChunkInt32(test.slice, test.size); len(test.expected) != len(got) {
			t.Errorf("In(%v, %d) = %v, expected: %v", test.slice, test.size, got, test.expected)
		}
	}
}

func TestChunkInt64(t *testing.T) {
	tests := []struct{
		slice []int64
		size int
		expected [][]int64
	}{
		{[]int64{1, 2, 3, 4, 5}, 2, [][]int64{{1, 2}, {3, 4}, {5}}},
		{[]int64{1, 2, 3}, 5, [][]int64{{1, 2, 3}}},
		{[]int64{1, 2}, 1, [][]int64{{1}, {2}}},
		{[]int64{1, 2}, 0, [][]int64{{1, 2}}},
	}

	for _, test := range tests {
		if got := ChunkInt64(test.slice, test.size); len(test.expected) != len(got) {
			t.Errorf("In(%v, %d) = %v, expected: %v", test.slice, test.size, got, test.expected)
		}
	}
}

func TestChunkStr(t *testing.T) {
	tests := []struct{
		slice []string
		size int
		expected [][]string
	}{
		{[]string{"1", "2", "3", "4", "5"}, 2, [][]string{{"1", "2"}, {"3", "4"}, {"5"}}},
		{[]string{"1", "2", "3"}, 5, [][]string{{"1", "2", "3"}}},
		{[]string{"1", "2"}, 1, [][]string{{"1"}, {"2"}}},
		{[]string{"1", "2"}, 0, [][]string{{"1", "2"}}},
	}

	for _, test := range tests {
		if got := ChunkStr(test.slice, test.size); len(test.expected) != len(got) {
			t.Errorf("In(%v, %d) = %v, expected: %v", test.slice, test.size, got, test.expected)
		}
	}
}
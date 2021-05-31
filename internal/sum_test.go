package internal

import (
	"testing"
)

func TestSum(t *testing.T) {
	var table = []struct {
		nums          []int
		expectedTotal int
	}{
		{
			nums:          []int{1, 2, 3, 4, 5},
			expectedTotal: 15,
		},
	}

	for _, row := range table {
		if total := Sum(row.nums...); total != row.expectedTotal {
			t.Errorf("Expected sum of '%d' rather than '%d'\n", row.expectedTotal, total)
		}
	}
}

package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectionSort(t *testing.T) {
	// ----- conditions ---------------------------------------------------------
	testCases := []struct {
		name           string
		input          []int
		expectedOutput []int
	}{
		{
			name:           "ReturnSortedElement",
			input:          []int{23, 0, 34, 1, 10},
			expectedOutput: []int{0, 1, 10, 23, 34},
		},
		{
			name:           "ReturnSortedElementWhenMinIsAtTheEnd",
			input:          []int{23, 2, 34, 1, 10, 0},
			expectedOutput: []int{0, 1, 2, 10, 23, 34},
		},
		{
			name:           "IdenticalElementReturnNoError",
			input:          []int{0, 0, 0},
			expectedOutput: []int{0, 0, 0},
		},
		{
			name:           "EmptyReturnNoError",
			input:          []int{},
			expectedOutput: []int{},
		},
	}

	for _, tc := range testCases {
		tc := tc // Capture range variable
		t.Run(tc.name, func(t *testing.T) {
			// ----- call -------------------------------------------------------
			selection_sort(tc.input, len(tc.input))
			// ----- verify -----------------------------------------------------
			assert.Equal(t, tc.expectedOutput, tc.input)
		})
	}
}

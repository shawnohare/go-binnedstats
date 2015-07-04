package binnedstats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHistogram(t *testing.T) {
	testCases := []struct {
		in      []float64
		numBins int
		out     []int
	}{
		// test case 1
		{
			in:      []float64{1, 2, 3},
			numBins: 2,
			out:     []int{1, 2},
		},
		// test case 2
		{
			in:      []float64{},
			numBins: 2,
			out:     []int{},
		},
	}

	for _, tt := range testCases {
		hist := Histogram(tt.in, tt.numBins)
		assert.Equal(t, tt.out, hist.Counts())
	}
}

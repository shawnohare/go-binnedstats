package binnedstats

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualWidthPartition(t *testing.T) {
	testCases := []struct {
		in      []float64
		numBins int
		out     [][]float64
	}{
		// test case 1
		{
			in:      []float64{1, 2, 3},
			numBins: 2,
			out:     [][]float64{[]float64{1}, []float64{2, 3}},
		},
		// test case 2
		{
			in:      []float64{},
			numBins: 2,
			out:     nil,
		},
	}

	for _, tt := range testCases {
		par := equalWidthPartition(tt.in, tt.numBins)
		assert.Equal(t, tt.out, par.Bins)
	}
}

package binnedstats

import "sort"

// Partition of data into subintervals (i.e., bins).
// A Partition structures data lying in some interval into
// appropriate sub-intervals.
type Partition struct {
	Bins  [][]float64
	Edges []float64
	Min   float64
	Max   float64
}

// A simple function to partition numberic data into equal width bins.
func equalWidthPartition(data []float64, numBins int) *Partition {
	n := len(data)
	// Deal with the empty input.
	if n == 0 || numBins < 1 {
		return &Partition{}
	}

	sort.Float64s(data)
	min := data[0]
	max := data[n-1]
	width := (max - min) / float64(numBins)

	// Deal with the case of identical data.
	if max == min {
		return &Partition{}
	}

	// Determine the data partition.
	rightEdges := make([]float64, numBins)
	for i := range rightEdges {
		rightEdges[i] = min + float64(i+1)*width
	}

	bins := make([][]float64, numBins)

	// Place each data point in the appropriate bin.
	for _, x := range data {
		var binIdx int
		for i, edge := range rightEdges {
			binIdx = i
			if x < edge {
				break
			}
		}
		bins[binIdx] = append(bins[binIdx], x)
	}

	return &Partition{
		Bins:  bins,
		Edges: rightEdges,
		Min:   min,
		Max:   max,
	}
}

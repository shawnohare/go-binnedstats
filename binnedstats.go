// Package binnedstats contains several structures and methods
// to facilitate the computation of statistics over partitions,
// such as histograms.
package binnedstats

// BinnedStats is a general purpose container representing a statistic
// function mapped across an arbitrary partition of the data.  When
// the statistic is just a count we have a histogram.
type BinnedStats struct {
	Bins []*Bin
}

func (bs *BinnedStats) Counts() []int {
	vals := make([]int, len(bs.Bins))
	for i, bin := range bs.Bins {
		vals[i] = bin.Count
	}
	return vals
}

func (bs *BinnedStats) Stats() []float64 {
	vals := make([]float64, len(bs.Bins))
	for i, bin := range bs.Bins {
		vals[i] = bin.Stat
	}
	return vals
}

func (bs *BinnedStats) Edges() []float64 {
	vals := make([]float64, len(bs.Bins))
	for i, bin := range bs.Bins {
		vals[i] = bin.RightEdge
	}
	return vals
}

func (bs *BinnedStats) Widths() []float64 {
	vals := make([]float64, len(bs.Bins))
	for i, bin := range bs.Bins {
		vals[i] = bin.Width
	}
	return vals
}

// A Bin containing statistics computed over a cell of some partition.
type Bin struct {
	LeftEdge  float64
	RightEdge float64
	Width     float64
	Count     int
	Stat      float64
}

// binnedStats computes the input statistic over each cell/bin of
// the input partition.
func binnedStats(par *Partition, stat func([]float64) float64) *BinnedStats {
	bins := par.Bins
	bs := &BinnedStats{
		Bins: make([]*Bin, len(bins)),
	}
	for i, bin := range bins {
		var leftEdge float64
		if i == 0 {
			leftEdge = par.Min
		} else {
			leftEdge = par.Edges[i-1]
		}
		bin := &Bin{
			Count:     len(bin),
			Stat:      stat(bin),
			LeftEdge:  leftEdge,
			RightEdge: par.Edges[i],
		}
		bin.Width = bin.RightEdge - bin.LeftEdge
		bs.Bins[i] = bin
	}
	return bs
}

func EqualWidthStats(data []float64, numBins int, stat func([]float64) float64) *BinnedStats {
	par := equalWidthPartition(data, numBins)
	return binnedStats(par, stat)
}

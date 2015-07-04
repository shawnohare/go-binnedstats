package binnedstats

func histogram(data []float64, numBins int) *BinnedStats {
	// Use length as the underlying statistic.
	f := func(xs []float64) float64 {
		return float64(len(xs))
	}

	return EqualWidthStats(data, numBins, f)
}

func Histogram(data []float64, numBins int) *BinnedStats {
	return histogram(data, numBins)
}

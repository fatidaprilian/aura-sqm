package filter

import "math"

func RejectOutlier(sample, baseline, maxRelativeDelta float64) bool {
	if baseline <= 0 || maxRelativeDelta <= 0 {
		return false
	}

	delta := math.Abs(sample - baseline)
	return delta/baseline > maxRelativeDelta
}

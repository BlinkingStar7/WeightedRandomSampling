package sampling

import (
	"math"
	"math/rand"
	"sort"
)

type WeightedRandomFast struct{}

func NewWeightedRandomFast() WeightedRandomSampling {
	return &WeightedRandomFast{}
}

// WRS selects k indices from weights without replacement, based on their probabilities.
func (w *WeightedRandomFast) WRS(weights []int, k int) []int {
	idxs := make([]int, len(weights))
	for i := range idxs {
		idxs[i] = i
	}

	// Fast Weighted Random Sampling
	// https://blog.taboola.com/going-old-school-designing-algorithms-fast-weighted-sampling-production/
	r := make([]float64, len(weights))
	for i, weight := range weights {
		r[i] = math.Pow(rand.Float64(), 1/float64(weight))
	}

	sort.Slice(idxs, func(i, j int) bool {
		return r[idxs[i]] > r[idxs[j]]
	})
	return idxs[:k]
}

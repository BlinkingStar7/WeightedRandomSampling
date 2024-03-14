package sampling

import (
	"math/rand"
)

type WeightedRandomNaive struct{}

func NewWeightedRandomNaive() weightedRandomSampling {
	return &WeightedRandomNaive{}
}

func (w *WeightedRandomNaive) selectOne(weights []int) int {
	sum := 0
	for _, weight := range weights {
		sum += weight
	}
	r := rand.Intn(sum)
	for i, w := range weights {
		r -= w
		if r < 0 {
			return i
		}
	}

	return -1
}

// WRS selects k indices from weights without replacement, based on their probabilities.
func (w *WeightedRandomNaive) WRS(weights []int, k int) []int {
	// Compute probabilities
	weightsCopy := make([]int, len(weights))
	copy(weightsCopy, weights)

	// Perform k selections based on probabilities
	selected := make([]int, 0, k)
	for len(selected) < k {
		idx := w.selectOne(weightsCopy)
		selected = append(selected, idx)
		weightsCopy[idx] = 0
	}

	return selected
}

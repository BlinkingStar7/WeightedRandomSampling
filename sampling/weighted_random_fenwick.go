package sampling

import (
	"math/rand"
)

type WeightedRandomFenwick struct {
	fw *FenwickTree
}

func NewWeightedRandomFenwick() weightedRandomSampling {
	return &WeightedRandomFenwick{}
}

func (w *WeightedRandomFenwick) selectOne() int {
	sum := w.fw.Sum(w.fw.size - 1)
	r := rand.Intn(sum)

	return w.fw.UpperBound(r)
}

// WRS selects k indices from weights without replacement, based on their probabilities.
func (w *WeightedRandomFenwick) WRS(weights []int, k int) []int {
	// Initializes a new Fenwick Tree
	w.fw = NewFenwickTree(weights)

	// Perform k selections based on probabilities
	selected := make([]int, 0, k)

	for len(selected) < k {
		idx := w.selectOne()
		selected = append(selected, idx)
		w.fw.Update(idx)
	}

	return selected
}

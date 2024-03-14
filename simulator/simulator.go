package simulator

import (
	"WeightedRandomSampling/sampling"
	"fmt"
	"math/rand"
	"time"
)

const (
	NUM_OF_TRIAL = 100000
)

type Simulator struct {
	Sampler sampling.WeightedRandomSampling
}

func NewSimulator(w sampling.WeightedRandomSampling) Simulator {
	return Simulator{
		Sampler: w,
	}
}

func (s *Simulator) Run() {
	probabilities := s.Simulate(NUM_OF_TRIAL)
	for _, p := range probabilities {
		fmt.Printf("%.5f\n", p)
	}
}

func (s *Simulator) Simulate(trials int) []float64 {
	N, K, weights := s.Input()

	rand.Seed(time.Now().UnixNano())
	countArr := make([]int, N)

	for i := 0; i < trials; i++ {
		result := s.Sampler.WRS(weights, K)
		for _, num := range result {
			countArr[num]++
		}
	}

	finalProb := make([]float64, N)
	for i, c := range countArr {
		finalProb[i] = float64(c) / float64(trials)
	}

	return finalProb
}

func (s *Simulator) Input() (int, int, []int) {
	var N, K int
	fmt.Scan(&N, &K)

	weights := make([]int, N)
	for i := range weights {
		fmt.Scan(&weights[i])
	}

	return N, K, weights
}

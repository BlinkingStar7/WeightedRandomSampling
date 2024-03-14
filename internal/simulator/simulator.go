package simulator

import (
	"WeightedRandomSampling/internal/sampling"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	NUM_OF_TRIAL = 100000
)

type Simulator struct {
	Sampler  sampling.WeightedRandomSampling
	FileName string
}

func NewSimulator(w sampling.WeightedRandomSampling, f string) Simulator {
	return Simulator{
		Sampler:  w,
		FileName: f,
	}
}

func (s *Simulator) Simulate() ([]float64, time.Duration) {
	N, K, weights, err := s.Input()

	if err != nil {
		fmt.Println(err)
		return nil, 0
	}

	countArr := make([]int, N)
	startTime := time.Now()

	for i := 0; i < NUM_OF_TRIAL; i++ {
		result := s.Sampler.WRS(weights, K)
		for _, num := range result {
			countArr[num]++
		}
	}
	elapsed := time.Since(startTime)

	finalProb := make([]float64, N)
	for i, c := range countArr {
		finalProb[i] = float64(c) / float64(NUM_OF_TRIAL) * 100
	}

	return finalProb, elapsed
}

func (s *Simulator) Input() (int, int, []int, error) {
	// read input from s.FileName
	file, err := os.Open(s.FileName)
	if err != nil {
		return 0, 0, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read first line for N and K
	scanner.Scan()
	firstLine := scanner.Text()
	parts := strings.Fields(firstLine)
	if len(parts) < 2 {
		return 0, 0, nil, fmt.Errorf("the first line should contain at least two numbers")
	}
	N, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, nil, err
	}
	K, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, nil, err
	}

	// Read the weights
	weights := make([]int, N)
	for i := 0; i < N; i++ {
		if !scanner.Scan() {
			return 0, 0, nil, fmt.Errorf("expected more weights")
		}
		weight, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return 0, 0, nil, err
		}
		weights[i] = weight
	}

	if err := scanner.Err(); err != nil {
		return 0, 0, nil, err
	}

	return N, K, weights, nil
}

package main

import (
	"WeightedRandomSampling/sampling"
	"WeightedRandomSampling/simulator"
)

func main() {
	NaiveSimulator := simulator.NewSimulator(sampling.NewWeightedRandomNaive())
	NaiveSimulator.Run()

	FenwickSimulator := simulator.NewSimulator(sampling.NewWeightedRandomFenwick())
	FenwickSimulator.Run()
}

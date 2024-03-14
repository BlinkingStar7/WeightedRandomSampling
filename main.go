package main

func main() {
	s := NewSimulator(NewWeightedRandomFenwick())
	s.Run()

	s = NewSimulator(NewWeightedRandomNaive())
	s.Run()
}

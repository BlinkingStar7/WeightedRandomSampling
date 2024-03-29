package main

import (
	"WeightedRandomSampling/internal/referenceSampling"
	"WeightedRandomSampling/internal/sampling"
	"WeightedRandomSampling/internal/simulator"
	"flag"
	"fmt"
	"math"
	"time"
)

const (
	NAIVE   = "naive"
	FENWICK = "fenwick"
	PYTHON  = "python"
	FAST    = "fast"
)

func runSimulation(c, fileName string, verbose bool) {
	var p []float64
	var t time.Duration
	switch c {
	case NAIVE:
		Simulator := simulator.NewSimulator(sampling.NewWeightedRandomNaive(), fileName)
		p, t = Simulator.Simulate()
	case FENWICK:
		Simulator := simulator.NewSimulator(sampling.NewWeightedRandomFenwick(), fileName)
		p, t = Simulator.Simulate()
	case FAST:
		Simulator := simulator.NewSimulator(sampling.NewWeightedRandomFast(), fileName)
		p, t = Simulator.Simulate()
	case PYTHON:
		p, t = referenceSampling.Simulate(fileName)
	default:
		fmt.Println("Invalid option for WhichVersionofWRS. Please use 'naive', 'fenwick', 'fast', or 'python'.")
	}

	fmt.Printf("%s Weighted Random Sampling Comparison on %d Trials\n", c, simulator.NUM_OF_TRIAL)
	fmt.Printf("Elapsed time: %v\n", t)
	if verbose {
		fmt.Println("Probability(%) of each element being selected:")
		for i, v := range p {
			fmt.Printf("Element %d: %.2f\n", i, v)
		}
	}
}

func main() {
	// Define command-line flags
	whichVersion := flag.String("wrs", "compare", "which version to run WRS (naive, fenwick, fast, python, compare)")
	fileName := flag.String("file-name", "invalid", "get a file name that is used")
	verbose := flag.Bool("verbose", false, "a bool")

	// Parse the flags
	flag.Parse()

	// Use the FileName flag value as needed for your simulator
	// For example, you might load scenarios from the file like this:

	if *fileName == "invalid" {
		fmt.Println("Invalid option for FileName. Please use a valid file name.")
		fmt.Println("Example: -file-name=scenarios/scenario1.txt")
		return
	}

	switch *whichVersion {
	case NAIVE:
		runSimulation(NAIVE, *fileName, *verbose)
	case FENWICK:
		runSimulation(FENWICK, *fileName, *verbose)
	case PYTHON:
		runSimulation(PYTHON, *fileName, *verbose)
	case FAST:
		runSimulation(FAST, *fileName, *verbose)
	case "both":
		FastSimulator := simulator.NewSimulator(sampling.NewWeightedRandomFast(), *fileName)
		FenwickSimulator := simulator.NewSimulator(sampling.NewWeightedRandomFenwick(), *fileName)

		pFast, tFast := FastSimulator.Simulate()
		pFenwick, tFenwick := FenwickSimulator.Simulate()

		// print table comparing probabilities
		fmt.Printf("Weighted Random Sampling Comparison on %d Trials\n", simulator.NUM_OF_TRIAL)
		fmt.Println("---------------------------------------------")

		var diffIdx []int
		for i := 0; i < len(pFast); i++ {
			// difference is max(fast, fenw) - min(fast, fenw)
			diff := math.Max(pFast[i], pFenwick[i]) - math.Min(pFast[i], pFenwick[i])
			if diff > 1 {
				diffIdx = append(diffIdx, i)
			}
		}

		fmt.Println("Difference is bigger than 1% p in the following indices:", diffIdx)
		fmt.Printf("Elapsed time: Fast: %v, Fenwick: %v\n\n", tFast, tFenwick)
		fmt.Println("---------------------------------------------")

		if *verbose {
			fmt.Println("Probability(%) of each element being selected:")
			fmt.Println("Element\t| Fast\t| Fenw\t| Diff(Max-Min)")
			fmt.Println("---------------------------------------------")
			for i := 0; i < len(pFast); i++ {
				// difference is max(fast, fenw) - min(fast, fenw)
				diff := math.Max(pFast[i], pFenwick[i]) - math.Min(pFast[i], pFenwick[i])
				if *verbose {
					fmt.Printf("%d\t| %.2f\t| %.2f\t| %.2f\n", i, pFast[i], pFenwick[i], diff)
				}
			}
		}
	case "compare":
		NaiveSimulator := simulator.NewSimulator(sampling.NewWeightedRandomNaive(), *fileName)
		FenwickSimulator := simulator.NewSimulator(sampling.NewWeightedRandomFenwick(), *fileName)

		pNaive, tNaive := NaiveSimulator.Simulate()
		pFenwick, tFenwick := FenwickSimulator.Simulate()
		pPython, tPython := referenceSampling.Simulate(*fileName)

		// print table comparing probabilities
		fmt.Printf("Weighted Random Sampling Comparison on %d Trials\n", simulator.NUM_OF_TRIAL)
		fmt.Println("---------------------------------------------")

		var diffIdx []int
		for i := 0; i < len(pNaive); i++ {
			// difference is max(naive, fenw, pyth) - min(naive, fenw, pyth)
			diff := math.Max(math.Max(pNaive[i], pFenwick[i]), pPython[i]) - math.Min(math.Min(pNaive[i], pFenwick[i]), pPython[i])
			if diff > 1 {
				diffIdx = append(diffIdx, i)
			}
		}

		fmt.Println("Difference is bigger than 1% p in the following indices:", diffIdx)
		fmt.Printf("Elapsed time: Naive: %v, Fenwick: %v, Python: %v\n\n", tNaive, tFenwick, tPython)
		fmt.Println("---------------------------------------------")

		if *verbose {
			fmt.Println("Probability(%) of each element being selected:")
			fmt.Println("Element\t| Naive\t| Fenw\t| Pyth\t| Diff(Max-Min)")
			fmt.Println("---------------------------------------------")
			for i := 0; i < len(pNaive); i++ {
				// difference is max(naive, fenw, pyth) - min(naive, fenw, pyth)
				diff := math.Max(math.Max(pNaive[i], pFenwick[i]), pPython[i]) - math.Min(math.Min(pNaive[i], pFenwick[i]), pPython[i])
				if *verbose {
					fmt.Printf("%d\t| %.2f\t| %.2f\t| %.2f\t| %.2f\n", i, pNaive[i], pFenwick[i], pPython[i], diff)
				}
			}
		}
	default:
		fmt.Println("Invalid option for WhichVersionofWRS. Please use 'naive', 'fenwick', or 'both'.")
	}
}

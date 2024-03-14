package referenceSampling

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func Simulate(fileName string) ([]float64, time.Duration) {
	cmd := exec.Command("python3", "./internal/referenceSampling/weighted_random_python.py", fileName)

	// run cmd and save result to variable result
	result, err := cmd.Output()

	if err != nil {
		fmt.Println("Error running Python script:", err)
		return nil, 0
	}

	// Split the result into lines
	lines := strings.Split(string(result), "\n")
	elapsedSeconds, err := strconv.ParseFloat(lines[0], 64)
	if err != nil {
		fmt.Println("Error converting Python elapsed time:", err)
		return nil, 0
	}

	// Convert float64 seconds to time.Duration
	elapsedDuration := time.Duration(elapsedSeconds * float64(time.Second))

	// Initialize a slice to hold the parsed floats
	var p []float64

	// Convert each line into a float64 and append it to the slice
	for _, line := range lines[1:] {
		if line == "" {
			continue // Skip empty lines
		}
		f, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Printf("Error converting line '%s' to float: %v\n", line, err)
			return nil, 0
		}
		p = append(p, f)
	}

	return p, elapsedDuration
}

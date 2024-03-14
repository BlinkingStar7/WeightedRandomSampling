package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const (
	WEIGHT_MAX = 10000
)

// generateScenario generates a scenario file based on user input
func generateScenario(N, K int) (string, error) {
	filePath := fmt.Sprintf("./scenarios/scenario_%d_%d.txt", N, K)

	// Open file for writing
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Use a buffered writer for efficient file writing
	writer := bufio.NewWriter(file)

	// Write the number of elements as the first line
	_, err = writer.WriteString(fmt.Sprintf("%d %d\n", N, K))
	if err != nil {
		return "", err
	}

	// Generate N random weights and write each to the file
	for i := 0; i < N; i++ {
		weight := rand.Intn(WEIGHT_MAX + 1) // Generates a number in [0, 9999]
		_, err = writer.WriteString(fmt.Sprintf("%d\n", weight))
		if err != nil {
			return "", err
		}
	}

	// Ensure all writes are flushed to the file
	if err := writer.Flush(); err != nil {
		return "", err
	}

	return fmt.Sprintf("Scenario file '%s' with %d elements created.", filePath, N), nil
}

func main() {
	// Example usage (replace with actual user input)
	var N, K int

	fmt.Print("input N: ")
	_, err := fmt.Scanf("%d", &N)
	if err != nil {
		fmt.Println("Error reading N:", err)
		return
	}

	fmt.Print("input K: ")
	_, err = fmt.Scanf("%d", &K)
	if err != nil {
		fmt.Println("Error reading K:", err)
		return
	}

	result, err := generateScenario(N, K)
	if err != nil {
		fmt.Println("Error generating scenario:", err)
	} else {
		fmt.Println(result)
	}
}

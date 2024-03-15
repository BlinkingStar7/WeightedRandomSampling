#!/bin/bash

# Directory containing scenario files
SCENARIOS_DIR="./scenarios"

# Directory to save results
RESULTS_DIR="./results/$(date +'%Y-%m-%d_%H-%M-%S')"

# Ensure the results directory exists
mkdir -p "$RESULTS_DIR"

# Loop through each file in the scenarios directory
for scenario_file in "$SCENARIOS_DIR"/*; do
    # Extract the base file name
    file_name=$(basename "$scenario_file")
    
    # Build the output file path
    result_file="$RESULTS_DIR/$file_name"
    
    # Run the Go command with the current scenario file
    go run cmd/compareWRS/main.go -file-name="$scenario_file" -wrs=compare -verbose=true > "$result_file"
    
    echo "Processed $file_name, results saved to $result_file"
done

echo "All scenarios processed."

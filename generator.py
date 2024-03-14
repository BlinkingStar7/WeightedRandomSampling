# Define a function to generate the scenario file based on user input
def generate_scenario(N, file_name):
    import random

    file_name = "./tests/" + file_name
    
    # Generate N random weights in the range 0 to 9999
    weights = [random.randint(0, 9999) for _ in range(N)]
    
    # Write the generated scenario to the specified file
    with open(file_name, 'w') as file:
        file.write(f"{N}\n")  # First line: number of elements
        for weight in weights:
            file.write(f"{weight}\n")  # Write each weight on a new line
            
    return f"Scenario file '{file_name}' with {N} elements created."

# Example usage (to be replaced with actual user input)
N = int(input("input N: "))
file_name = input("input FileName: ")
print(generate_scenario(N, file_name))


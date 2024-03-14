import numpy as np
import sys
import time

TRIAL = 100000

def WRS(weights, k):
    total = sum(weights)
    prob = [w/total for w in weights]
    population = list(range(len(weights)))
    return np.random.choice(a=population, size=k, replace=False, p=prob)

def Simulate(f):
    # read input from f
    n, k = map(int, f.readline().split())
    weights = [int(f.readline()) for _ in range(n)]

    total = [0] * n
    # call the WRS function
    start = time.time()
    for i in range(TRIAL):
        result = WRS(weights, k)
        for j in result:
            total[j] += 1

    elapsed = time.time() - start

    # return the result
    return [t/TRIAL*100 for t in total], elapsed

def main(argv):
    # open the input file that from command line argument
    with open(argv[1], "r") as f:
        # call the Simulate function
        result, t = Simulate(f)
        print(t)
        # print the result
        for r in result:
            print(r)

if __name__ == "__main__":
    main(sys.argv)
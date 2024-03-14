import random
import numpy as np

def WRS(weights, k=1):
    total = sum(weights)
    prob = [w/total for w in weights]
    population = list(range(len(weights)))
    return np.random.choice(a=population, size=k, replace=False, p=prob)

# get lnput
N, K = map(int, input().split())
weights = list(map(int, input().split()))
trial = 100*100
count = [0] * N

# trial
for _ in range(trial):
    result = WRS(weights, K)
    for num in result:
        count[num] += 1

# print output
finalProb = [c / trial for c in count]
for p in finalProb:
    print("%.2f"%p)


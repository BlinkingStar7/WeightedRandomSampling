# Description
원소의 개수 N, 뽑고자 하는 개수 K, 각 원소의 가중치 Weights가 주어질 때 Weighted Random Sampling(WRS)를 다음과 같은 3가지 방식으로 구현하고 정확도와 성능을 비교합니다.

## Sampling 1. Python Reference Solution
- Python의 `np.random.choice` 를 통해 구현된 레퍼런스 WRS 입니다.
- https://github.com/BlinkingStar7/WeightedRandomSampling/blob/d5d65fe406efdf6e6a34805294b80eab90f59d54/internal/referenceSampling/weighted_random_python.py#L1-L42

## Sampling 2. Golang Naive Approach (O(N^2))
- [버즈빌 기술 블로그](https://tech.buzzvil.com/blog/weighted-random-shuffling/)에서 제안한 Naive한 구현입니다.
- https://github.com/BlinkingStar7/WeightedRandomSampling/blob/d5d65fe406efdf6e6a34805294b80eab90f59d54/internal/sampling/weighted_random_naive.go#L1-L44

## Sampling 3. Golang Fenwick Tree Approach (O(NlogN))
- [버즈빌 기술 블로그](https://tech.buzzvil.com/blog/weighted-random-shuffling/)에서 제안한 BST를 이용한 방식의 메모리 사용량을 개선시킨 구현입니다.
- https://github.com/BlinkingStar7/WeightedRandomSampling/blob/d5d65fe406efdf6e6a34805294b80eab90f59d54/internal/sampling/weighted_random_fenwick.go#L1-L37

# Scenario
N, K를 입력으로 받아서 무작위 Weights를 가지는 scenario 파일을 생성합니다. 생성기는 [cmd/scenarioGenerator](https://github.com/BlinkingStar7/WeightedRandomSampling/blob/d5d65fe406efdf6e6a34805294b80eab90f59d54/cmd/scenarioGenerator/main.go#L1-L75)에 위치하여 있습니다.
생성된 시나리오 파일은 `./scenario/scenario_{N}_{K}.txt`에 생성됩니다.

## Example
```
go run cmd/scenarioGenerator/main.go
```

# Simulation
주어진 시나리오 파일에 대하여 `NUM_OF_TRIAL (Default: 100,000회)`만큼의 WRS를 수행합니다.

## Parameter
- `-file-name`: 실행시키고자 하는 시나리오 이름.
- `wrs`: 어떤 sampling 방법을 사용할지. { naive | fenwick | python | compare } 중 택 1. (Default: copmare)
  - compare 모드는 3가지 sampling 방법을 모두 실행하여서 정확도와 성능을 비교합니다.
- `verboase`: 각 element들이 뽑힐 확률을 출력할지 여부입니다. (Default: false)

## Example
```
go run cmd/compareWRS/main.go -file-name=scenarios/scenario_10_1.txt -wrs=compare -verbose=true
```

# Result
기본으로 제공된 시나리오 파일들에 대하여 분석 결과를 `/result`폴더에 추가하였습니다.
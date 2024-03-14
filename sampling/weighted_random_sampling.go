package sampling

type weightedRandomSampling interface {
	WRS(weights []int, k int) []int
}

type DefaultWeightedRandomSampling struct{}

func NewDefaultWeightedRandomSampling() weightedRandomSampling {
	return &DefaultWeightedRandomSampling{}
}

func (d *DefaultWeightedRandomSampling) WRS(weights []int, k int) []int {
	return nil
}

package sampling

type WeightedRandomSampling interface {
	WRS(weights []int, k int) []int
}

type DefaultWeightedRandomSampling struct{}

func NewDefaultWeightedRandomSampling() WeightedRandomSampling {
	return &DefaultWeightedRandomSampling{}
}

func (d *DefaultWeightedRandomSampling) WRS(weights []int, k int) []int {
	return nil
}

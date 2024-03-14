package sampling

type FenwickTree struct {
	tree []int
	size int
}

// NewFenwickTree creates a new Fenwick Tree of given weights
func NewFenwickTree(weights []int) *FenwickTree {
	_size := len(weights)
	_tree := make([]int, _size+1)

	for i := 1; i <= _size; i++ {
		_tree[i] += weights[i-1]
		j := i + (i & -i)
		if j <= _size {
			_tree[j] += _tree[i]
		}
	}

	return &FenwickTree{
		tree: _tree,
		size: _size,
	}
}

// Update set value to 0 at index i.
func (fw *FenwickTree) Update(i int) {
	val := fw.Sum(i) - fw.Sum(i-1)
	i++ // Fenwick Tree indices start from 1
	for i <= fw.size {
		fw.tree[i] -= val
		i += i & -i // Move to the next index
	}
}

// Sum returns the prefix sum from start to i.
func (fw *FenwickTree) Sum(i int) int {
	i++ // Adjust index to match internal representation
	sum := 0
	for i > 0 {
		sum += fw.tree[i]
		i -= i & -i // Move to parent index
	}
	return sum
}

func (fw *FenwickTree) GetTotalSum() int {
	return fw.Sum(fw.size - 1)
}

// UpperBound returns the index of the first element in the tree that has a prefix sum greater than bound.
func (fw *FenwickTree) UpperBound(bound int) int {
	idx := 0
	sum := 0
	k := 1
	for k < fw.size {
		k <<= 1
	}

	for ; k > 0; k >>= 1 {
		if idx+k <= fw.size && (fw.tree[idx+k] == 0 || sum+fw.tree[idx+k] < bound) {
			sum += fw.tree[idx+k]
			idx += k
		}
	}

	return idx
}

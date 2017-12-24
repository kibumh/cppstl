package algorithm

// Getter can get a value at 'i'th position.
type Getter interface {
	Get(i int) interface{}
}

// Lenner has a length
type Lenner interface {
	Len() int
}

// Swapper can swap
type Swapper interface {
	Swap(i, j int)
}

// LenSwapper is both a Lenner and a Swapper
type LenSwapper interface {
	Lenner
	Swapper
}

// GetLenner is both a Getter and a Lenner
type GetLenner interface {
	Getter
	Lenner
}

// GetSwapper is both a Getter and a Swapper
type GetSwapper interface {
	Swapper
	Getter
}

// GetLenSwapper is a Getter, a Lenner, and a Swapper
type GetLenSwapper interface {
	Getter
	Lenner
	Swapper
}

// ReverseRange reverse a given container within a range [begin, end)
func ReverseRange(s Swapper, begin, end int) {
	mid := (begin + end) / 2
	for i := begin; i < mid; i++ {
		s.Swap(begin+i, end-1-i)
	}
}

// Reverse reverse a given container.
func Reverse(ls LenSwapper) {
	ReverseRange(ls, 0, ls.Len())
}

// RotateRange rotates elements in a given container within a range [begin, end)
// It returns a index to the value previously at 'begin'.
func RotateRange(s Swapper, begin, middle, end int) int {
	if begin > middle || middle > end {
		return begin
	}

	retidx := end - middle + begin
	next := middle
	for begin != next {
		s.Swap(begin, next)
		begin++
		next++
		if next == end {
			next = middle
		}
		if begin == middle {
			middle = next
		}
	}
	return retidx
}

// Rotate rotate elements in such a way that the value at 'middle' becomes the first element.
func Rotate(ls LenSwapper, middle int) int {
	return RotateRange(ls, 0, middle, ls.Len())
}

// StablePartitionRange partitions in two groups.
func StablePartitionRange(gs GetSwapper, begin, end int, pred func(v interface{}) bool) int {
	if len := end - begin; len == 0 {
		return begin
	} else if len == 1 {
		if pred(gs.Get(begin)) {
			return begin + 1
		}
		return begin
	} else {
		middle := (begin + end) / 2
		return RotateRange(gs,
			StablePartitionRange(gs, begin, middle, pred),
			middle,
			StablePartitionRange(gs, middle, end, pred))
	}
}

// StablePartition partitions in two groups.
func StablePartition(gls GetLenSwapper, pred func(v interface{}) bool) int {
	return StablePartitionRange(gls, 0, gls.Len(), pred)
}

// AllOfRange returns true only if all elements meet a given condition
func AllOfRange(g Getter, begin, end int, pred func(v interface{}) bool) bool {
	for i := begin; i < end; i++ {
		if !pred(g.Get(i)) {
			return false
		}
	}
	return true
}

// AllOf returns true only if all elements meet a given condition
func AllOf(gl GetLenner, pred func(v interface{}) bool) bool {
	return AllOfRange(gl, 0, gl.Len(), pred)
}

// NoneOfRange returns true only if no element meets a given condition
func NoneOfRange(g Getter, begin, end int, pred func(v interface{}) bool) bool {
	for i := begin; i < end; i++ {
		if pred(g.Get(i)) {
			return false
		}
	}
	return true
}

// NoneOf returns true only if no element meets a given condition
func NoneOf(gl GetLenner, pred func(v interface{}) bool) bool {
	return NoneOfRange(gl, 0, gl.Len(), pred)
}

// AnyOfRange returns true if any element meet a given condition
func AnyOfRange(g Getter, begin, end int, pred func(v interface{}) bool) bool {
	for i := begin; i < end; i++ {
		if pred(g.Get(i)) {
			return true
		}
	}
	return false
}

// AnyOf returns true if any element meet a given condition
func AnyOf(gl GetLenner, pred func(v interface{}) bool) bool {
	return AnyOfRange(gl, 0, gl.Len(), pred)
}

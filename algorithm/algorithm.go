// package algorithm implements some useful algorithms in C++ STL.
//
// API signatures is a copy of sort package in stdlib.
//   - SomeAlgorithm : receives a predefined interface. (sort.Sort)
//   - SomeAlgorithmRange : receives a predefined interface and a range [begin, end).
//   - SomeAlgorithmSlice : receives a slice. (sort.Slice)

package algorithm

import (
	"reflect"
)

// Getter can get a value at 'i'th position.
type Getter interface {
	Get(i int) interface{}
}

// Lesser compares two elements.
type Lesser interface {
	Less(i, j int) bool
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

// LessSwapper is both a Lesser and a Swapper
type LessSwapper interface {
	Lesser
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

// LenLessSwapper is both a Lenner, a Lesser and a Swapper
type LenLessSwapper interface {
	Lenner
	Lesser
	Swapper
}

// ReverseRange reverse a given container within a range [begin, end).
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

// ReverseSlice is a slice version of Reverse.
func ReverseSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()

	mid := length / 2
	for i := 0; i < mid; i++ {
		swap(i, length-1-i)
	}
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

// RotateSlice is a Rotate function with a slice.
func RotateSlice(slice interface{}, middle int) int {
	return rotateSliceImpl(reflect.Swapper(slice), 0, middle, reflect.ValueOf(slice).Len())
}

func rotateSliceImpl(swap func(i, j int), begin, middle, end int) int {
	if begin > middle || middle > end {
		return begin
	}

	retidx := end - middle + begin
	next := middle
	for begin != next {
		swap(begin, next)
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

func stablePartitionSliceImpl(swap func(i, j int), begin, end int, pred func(i int) bool) int {
	if len := end - begin; len == 0 {
		return begin
	} else if len == 1 {
		if pred(begin) {
			return begin + 1
		}
		return begin
	} else {
		middle := (begin + end) / 2
		return rotateSliceImpl(swap,
			stablePartitionSliceImpl(swap, begin, middle, pred),
			middle,
			stablePartitionSliceImpl(swap, middle, end, pred))
	}
}

// StablePartitionSlice is a Rotate function with a slice.
func StablePartitionSlice(slice interface{}, pred func(i int) bool) int {
	return stablePartitionSliceImpl(reflect.Swapper(slice), 0, reflect.ValueOf(slice).Len(), pred)
}

// AllOfRange returns true only if all elements meet a given condition.
func AllOfRange(g Getter, begin, end int, pred func(v interface{}) bool) bool {
	for i := begin; i < end; i++ {
		if !pred(g.Get(i)) {
			return false
		}
	}
	return true
}

// AllOf returns true only if all elements meet a given condition.
func AllOf(gl GetLenner, pred func(v interface{}) bool) bool {
	return AllOfRange(gl, 0, gl.Len(), pred)
}

// AllOfSlice returns true only if all elements in a given slice meet a given condition.
func AllOfSlice(slice interface{}, pred func(i int) bool) bool {
	length := reflect.ValueOf(slice).Len()

	for i := 0; i < length; i++ {
		if !pred(i) {
			return false
		}
	}
	return true
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

// NoneOfSlice returns true only if no element in a given slice meets a given condition.
func NoneOfSlice(slice interface{}, pred func(i int) bool) bool {
	length := reflect.ValueOf(slice).Len()

	for i := 0; i < length; i++ {
		if pred(i) {
			return false
		}
	}
	return true
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

// AnyOfSlice returns true if any element in a given slice meets a given condition.
func AnyOfSlice(slice interface{}, pred func(i int) bool) bool {
	length := reflect.ValueOf(slice).Len()

	for i := 0; i < length; i++ {
		if pred(i) {
			return true
		}
	}
	return false
}

// NthElementRange rearranges a range [begin, end) in such a way that the element at nth(k) position is the element that would occur in that position if a range is sorted. All of the other elements in a range before nth position is less than or equal to the new nth element.
func NthElementRange(ls LessSwapper, begin, end, k int) {
	nthElementSliceImpl(ls.Swap, ls.Less, begin, end, k)
}

// NthElement rearranges a slice in such a way that the element at nth(k) position is the element that would occur in that position if slice is sorted. All of the other elements before nth position is less than or equal to the new nth element.
func NthElement(lls LenLessSwapper, k int) {
	NthElementRange(lls, 0, lls.Len(), k)
}

// NthElementSlice rearranges a slice in such a way that the element at nth(k) position is the element that would occur in that position if slice is sorted. All of the other elements before nth position is less than or equal to the new nth element.
func NthElementSlice(slice interface{}, less func(i, j int) bool, k int) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()

	nthElementSliceImpl(swap, less, 0, length, k)
}

func nthElementSliceImpl(swap func(i, j int), less func(i, j int) bool, begin, end, k int) {
	if begin+1 >= end {
		return
	}

	pidx := begin
	pidx = partitionSliceImpl(swap, less, begin, end, pidx)
	if k == pidx {
		return
	} else if k < pidx {
		nthElementSliceImpl(swap, less, begin, pidx, k)
		return
	} else {
		nthElementSliceImpl(swap, less, pidx+1, end, k)
		return
	}
}

func partitionSliceImpl(swap func(i, j int), less func(i, j int) bool, begin, end, pidx int) int {
	swap(end-1, pidx)
	pidx = end - 1
	sidx := begin
	for i := begin; i < end-1; i++ {
		if less(i, pidx) {
			swap(sidx, i)
			sidx++
		}
	}
	swap(end-1, sidx)
	return sidx
}

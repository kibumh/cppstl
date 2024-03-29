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

func transformPred(pred func(v interface{}) bool, g Getter) func(i int) bool {
	return func(i int) bool {
		return pred(g.Get(i))
	}
}

// Reverse reverse a given container.
func Reverse(ls LenSwapper) {
	ReverseRange(ls, 0, ls.Len())
}

// ReverseRange reverse a given container within a range [begin, end).
func ReverseRange(s Swapper, begin, end int) {
	reverseImpl(s.Swap, begin, end)
}

// ReverseSlice is a slice version of Reverse.
func ReverseSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	swap := reflect.Swapper(slice)
	length := rv.Len()
	reverseImpl(swap, 0, length)
}

func reverseImpl(swap func(i, j int), begin, end int) {
	mid := (begin + end) / 2
	for i := 0; i < mid-begin; i++ {
		swap(begin+i, end-1-i)
	}
}

// Rotate rotate elements in such a way that the value at 'middle' becomes the first element.
func Rotate(ls LenSwapper, middle int) int {
	return RotateRange(ls, 0, middle, ls.Len())
}

// RotateRange rotates elements in a given container within a range [begin, end)
// It returns a index to the value previously at 'begin'.
func RotateRange(s Swapper, begin, middle, end int) int {
	return rotateImpl(s.Swap, begin, middle, end)
}

// RotateSlice is a Rotate function with a slice.
func RotateSlice(slice interface{}, middle int) int {
	return rotateImpl(reflect.Swapper(slice), 0, middle, reflect.ValueOf(slice).Len())
}

func rotateImpl(swap func(i, j int), begin, middle, end int) int {
	if begin > middle || middle > end {
		return begin
	}
	reverseImpl(swap, begin, middle)
	reverseImpl(swap, middle, end)
	reverseImpl(swap, begin, end)
	return end - middle + begin
}

// StablePartition partitions in two groups.
func StablePartition(gls GetLenSwapper, pred func(v interface{}) bool) int {
	return StablePartitionRange(gls, 0, gls.Len(), pred)
}

// StablePartitionRange partitions in two groups.
func StablePartitionRange(gs GetSwapper, begin, end int, pred func(v interface{}) bool) int {
	return stablePartitionSliceImpl(gs.Swap, begin, end, transformPred(pred, gs))
}

// StablePartitionSlice is a Rotate function with a slice.
func StablePartitionSlice(slice interface{}, pred func(i int) bool) int {
	return stablePartitionSliceImpl(reflect.Swapper(slice), 0, reflect.ValueOf(slice).Len(), pred)
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
		return rotateImpl(swap,
			stablePartitionSliceImpl(swap, begin, middle, pred),
			middle,
			stablePartitionSliceImpl(swap, middle, end, pred))
	}
}

// AllOf returns true only if all elements meet a given condition.
func AllOf(gl GetLenner, pred func(v interface{}) bool) bool {
	return AllOfRange(gl, 0, gl.Len(), pred)
}

// AllOfRange returns true only if all elements meet a given condition.
func AllOfRange(g Getter, begin, end int, pred func(v interface{}) bool) bool {
	return allOfImpl(begin, end, transformPred(pred, g))
}

// AllOfSlice returns true only if all elements in a given slice meet a given condition.
func AllOfSlice(slice interface{}, pred func(i int) bool) bool {
	return allOfImpl(0, reflect.ValueOf(slice).Len(), pred)
}

func allOfImpl(begin, end int, pred func(i int) bool) bool {
	for i := begin; i < end; i++ {
		if !pred(i) {
			return false
		}
	}
	return true
}

// NoneOf returns true only if no element meets a given condition
func NoneOf(gl GetLenner, pred func(v interface{}) bool) bool {
	return NoneOfRange(gl, 0, gl.Len(), pred)
}

// NoneOfRange returns true only if no element meets a given condition
func NoneOfRange(g Getter, begin, end int, pred func(v interface{}) bool) bool {
	return noneOfImpl(begin, end, transformPred(pred, g))
}

// NoneOfSlice returns true only if no element in a given slice meets a given condition.
func NoneOfSlice(slice interface{}, pred func(i int) bool) bool {
	return noneOfImpl(0, reflect.ValueOf(slice).Len(), pred)
}

func noneOfImpl(begin, end int, pred func(i int) bool) bool {
	for i := begin; i < end; i++ {
		if pred(i) {
			return false
		}
	}
	return true
}

// AnyOf returns true if any element meet a given condition
func AnyOf(gl GetLenner, pred func(v interface{}) bool) bool {
	return AnyOfRange(gl, 0, gl.Len(), pred)
}

// AnyOfRange returns true if any element meet a given condition
func AnyOfRange(g Getter, begin, end int, pred func(v interface{}) bool) bool {
	return anyOfImpl(begin, end, transformPred(pred, g))
}

// AnyOfSlice returns true if any element in a given slice meets a given condition.
func AnyOfSlice(slice interface{}, pred func(i int) bool) bool {
	return anyOfImpl(0, reflect.ValueOf(slice).Len(), pred)
}

func anyOfImpl(begin, end int, pred func(i int) bool) bool {
	for i := begin; i < end; i++ {
		if pred(i) {
			return true
		}
	}
	return false
}

// NthElement rearranges a slice in such a way that the element at nth(k) position is
// the element that would occur in that position if slice is sorted.
// All of the other elements before nth position is less than or equal to the new nth element.
func NthElement(lls LenLessSwapper, k int) {
	NthElementRange(lls, 0, lls.Len(), k)
}

// NthElementRange rearranges a range [begin, end) in such a way that the element at
// nth(k) position is the element that would occur in that position if a range is sorted.
// All of the other elements in a range before nth position is less than or equal to the new nth element.
func NthElementRange(ls LessSwapper, begin, end, k int) {
	nthElementSliceImpl(ls.Swap, ls.Less, begin, end, k)
}

// NthElementSlice rearranges a slice in such a way that the element at nth(k) position is
// the element that would occur in that position if slice is sorted.
// All of the other elements before nth position is less than or equal to the new nth element.
func NthElementSlice(slice interface{}, less func(i, j int) bool, k int) {
	nthElementSliceImpl(reflect.Swapper(slice), less, 0, reflect.ValueOf(slice).Len(), k)
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

package algorithm_test

import (
	"fmt"

	"github.com/kibumh/cppstl/algorithm"
)

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}

func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is intSlice) Get(i int) interface{} {
	return is[i]
}

func ExampleReverse() {
	is := intSlice{1, 2, 3, 4, 5}
	fmt.Println(is)

	algorithm.Reverse(is)
	fmt.Println(is)

	isInner := is[1:3]
	algorithm.Reverse(isInner)
	fmt.Println(is)
	// Output:
	// [1 2 3 4 5]
	// [5 4 3 2 1]
	// [5 3 4 2 1]
}

func ExampleReverseRange() {
	is := intSlice{1, 2, 3, 4, 5}
	fmt.Println(is)

	algorithm.ReverseRange(is, 0, 5)
	fmt.Println(is)

	algorithm.ReverseRange(is, 1, 3)
	fmt.Println(is)
	// Output:
	// [1 2 3 4 5]
	// [5 4 3 2 1]
	// [5 3 4 2 1]
}

func ExampleReverseSlice() {
	is := []int{1, 2, 3, 4, 5}
	fmt.Println(is)

	algorithm.ReverseSlice(is)
	fmt.Println(is)

	isInner := is[1:3]
	algorithm.ReverseSlice(isInner)
	fmt.Println(is)
	// Output:
	// [1 2 3 4 5]
	// [5 4 3 2 1]
	// [5 3 4 2 1]
}

func ExampleRotate() {
	is := intSlice{1, 2, 3, 4, 5}
	fmt.Println(is)

	pos := algorithm.Rotate(is, 2)
	fmt.Println(is, pos)

	isInner := is[:3]
	pos = algorithm.Rotate(isInner, 1)
	fmt.Println(is, pos)

	// Output:
	// [1 2 3 4 5]
	// [3 4 5 1 2] 3
	// [4 5 3 1 2] 2
}

func ExampleRotateRange() {
	is := intSlice{1, 2, 3, 4, 5}
	fmt.Println(is)

	pos := algorithm.RotateRange(is, 0, 2, 5)
	fmt.Println(is, pos)

	pos = algorithm.RotateRange(is, 0, 1, 3)
	fmt.Println(is, pos)

	// Output:
	// [1 2 3 4 5]
	// [3 4 5 1 2] 3
	// [4 5 3 1 2] 2
}

func ExampleRotateSlice() {
	is := []int{1, 2, 3, 4, 5}
	fmt.Println(is)

	pos := algorithm.RotateSlice(is, 2)
	fmt.Println(is, pos)

	isInner := is[:3]
	pos = algorithm.RotateSlice(isInner, 1)
	fmt.Println(is, pos)

	// Output:
	// [1 2 3 4 5]
	// [3 4 5 1 2] 3
	// [4 5 3 1 2] 2
}

func ExampleStablePartition() {
	is := intSlice{1, 2, 3, 4, 5}
	fmt.Println(is)

	pred := func(v interface{}) bool {
		i := v.(int)
		return i%2 == 0
	}

	pos := algorithm.StablePartition(is, pred)
	fmt.Println(is, pos)

	// Output:
	// [1 2 3 4 5]
	// [2 4 1 3 5] 2
}

func ExampleStablePartitionRange() {
	is := intSlice{1, 2, 3, 4, 5}
	fmt.Println(is)

	pred := func(v interface{}) bool {
		i := v.(int)
		return i%2 == 0
	}

	pos := algorithm.StablePartitionRange(is, 0, 5, pred)
	fmt.Println(is, pos)

	// Output:
	// [1 2 3 4 5]
	// [2 4 1 3 5] 2
}

func ExampleStablePartitionSlice() {
	is := []int{1, 2, 3, 4, 5}
	fmt.Println(is)

	pred := func(i int) bool {
		return is[i]%2 == 0
	}

	pos := algorithm.StablePartitionSlice(is, pred)
	fmt.Println(is, pos)

	// Output:
	// [1 2 3 4 5]
	// [2 4 1 3 5] 2
}

func ExampleAllOf() {
	pred := func(v interface{}) bool {
		return v.(int)%2 == 0
	}
	fmt.Println(algorithm.AllOf(intSlice{1, 2, 3, 4, 5}, pred))
	fmt.Println(algorithm.AllOf(intSlice{2, 4, 6, 8, 10}, pred))

	// Output:
	// false
	// true
}

func ExampleNoneOf() {
	pred := func(v interface{}) bool {
		return v.(int)%2 == 0
	}
	fmt.Println(algorithm.NoneOf(intSlice{1, 3, 5, 7, 9}, pred))
	fmt.Println(algorithm.NoneOf(intSlice{2, 4, 6, 8, 10}, pred))

	// Output:
	// true
	// false
}

func ExampleAnyOf() {
	pred := func(v interface{}) bool {
		return v.(int)%2 != 0
	}
	fmt.Println(algorithm.AnyOf(intSlice{1, 2, 3, 4, 5}, pred))
	fmt.Println(algorithm.AnyOf(intSlice{2, 4, 6, 8, 10}, pred))

	// Output:
	// true
	// false
}

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

func isEven(v interface{}) bool {
	return v.(int)%2 == 0
}

func isEvenFn(is []int) func(int) bool {
	return func(i int) bool {
		return is[i]%2 == 0
	}
}

func isLessFn(is []int) func(i, j int) bool {
	return func(i, j int) bool {
		return is[i] < is[j]
	}
}

func ExampleReverse() {
	is := intSlice{1, 2, 3, 4, 5}

	algorithm.Reverse(is)
	fmt.Println(is)

	algorithm.Reverse(is[1:3])
	fmt.Println(is)

	// Output:
	// [5 4 3 2 1]
	// [5 3 4 2 1]
}

func ExampleReverseRange() {
	is := intSlice{1, 2, 3, 4, 5}

	algorithm.ReverseRange(is, 0, 5)
	fmt.Println(is)

	algorithm.ReverseRange(is, 1, 3)
	fmt.Println(is)

	// Output:
	// [5 4 3 2 1]
	// [5 3 4 2 1]
}

func ExampleReverseSlice() {
	is := []int{1, 2, 3, 4, 5}

	algorithm.ReverseSlice(is)
	fmt.Println(is)

	algorithm.ReverseSlice(is[1:3])
	fmt.Println(is)

	// Output:
	// [5 4 3 2 1]
	// [5 3 4 2 1]
}

func ExampleRotate() {
	is := intSlice{1, 2, 3, 4, 5}

	pos := algorithm.Rotate(is, 2)
	fmt.Println(is, pos)

	pos = algorithm.Rotate(is[:3], 1)
	fmt.Println(is, pos)

	// Output:
	// [3 4 5 1 2] 3
	// [4 5 3 1 2] 2
}

func ExampleRotateRange() {
	is := intSlice{1, 2, 3, 4, 5}

	pos := algorithm.RotateRange(is, 0, 2, 5)
	fmt.Println(is, pos)

	pos = algorithm.RotateRange(is, 0, 1, 3)
	fmt.Println(is, pos)

	// Output:
	// [3 4 5 1 2] 3
	// [4 5 3 1 2] 2
}

func ExampleRotateSlice() {
	is := []int{1, 2, 3, 4, 5}

	pos := algorithm.RotateSlice(is, 2)
	fmt.Println(is, pos)

	pos = algorithm.RotateSlice(is[:3], 1)
	fmt.Println(is, pos)

	// Output:
	// [3 4 5 1 2] 3
	// [4 5 3 1 2] 2
}

func ExampleStablePartition() {
	is := intSlice{1, 2, 3, 4, 5}

	pos := algorithm.StablePartition(is, isEven)
	fmt.Println(is, pos)

	// Output:
	// [2 4 1 3 5] 2
}

func ExampleStablePartitionRange() {
	is := intSlice{1, 2, 3, 4, 5}

	pos := algorithm.StablePartitionRange(is, 0, 5, isEven)
	fmt.Println(is, pos)

	// Output:
	// [2 4 1 3 5] 2
}

func ExampleStablePartitionSlice() {
	is := []int{1, 2, 3, 4, 5}

	pos := algorithm.StablePartitionSlice(is, isEvenFn(is))
	fmt.Println(is, pos)

	// Output:
	// [2 4 1 3 5] 2
}

func ExampleAllOf() {
	is := intSlice{1, 2, 3, 4, 5}
	fmt.Println(algorithm.AllOf(is, isEven))

	is2 := intSlice{2, 4, 6, 8, 10}
	fmt.Println(algorithm.AllOf(is2, isEven))

	// Output:
	// false
	// true
}

func ExampleAllOfSlice() {
	is := []int{1, 2, 3, 4, 5}
	fmt.Println(algorithm.AllOfSlice(is, isEvenFn(is)))

	is2 := []int{2, 4, 6, 8, 10}
	fmt.Println(algorithm.AllOfSlice(is2, isEvenFn(is2)))

	// Output:
	// false
	// true
}

func ExampleNoneOf() {
	is := intSlice{1, 2, 3, 4, 5}
	fmt.Println(algorithm.NoneOf(is, isEven))

	is2 := intSlice{1, 3, 5, 7, 9}
	fmt.Println(algorithm.NoneOf(is2, isEven))

	// Output:
	// false
	// true
}

func ExampleNoneOfSlice() {
	is := []int{1, 2, 3, 4, 5}
	fmt.Println(algorithm.NoneOfSlice(is, isEvenFn(is)))

	is2 := []int{1, 3, 5, 7, 9}
	fmt.Println(algorithm.NoneOfSlice(is2, isEvenFn(is2)))

	// Output:
	// false
	// true
}

func ExampleAnyOf() {
	is := intSlice{1, 2, 3, 4, 5}
	fmt.Println(algorithm.AnyOf(is, isEven))

	is2 := intSlice{1, 3, 5, 7, 9}
	fmt.Println(algorithm.AnyOf(is2, isEven))

	// Output:
	// true
	// false
}

func ExampleAnyOfSlice() {
	is := []int{1, 2, 3, 4, 5}
	fmt.Println(algorithm.AnyOfSlice(is, isEvenFn(is)))

	is2 := []int{1, 3, 5, 7, 9}
	fmt.Println(algorithm.AnyOfSlice(is2, isEvenFn(is2)))

	// Output:
	// true
	// false
}

func ExampleNthElementSlice() {
	is := []int{3, 1, 4, 5, 2}

	for k := 0; k < 5; k++ {
		newis := append([]int(nil), is...)
		algorithm.NthElementSlice(newis, isLessFn(newis), k)
		fmt.Println(newis[k])
	}

	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
}

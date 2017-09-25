package heap

import (
	"testing"
	utils "common/testUtils"
)

func TestHeapSort(t *testing.T) {
	val := []int{}
	HeapSort(val)

	val = []int{12, 2, 11, 9, 34}
	HeapSort(val)
	expected := []int{2, 9, 11, 12, 34}
	utils.AssertArrayDeepEqual(t, val, expected)

	val = []int{33, 12, 2, 11, 9, 34, 7,}
	HeapSort(val)
	expected = []int{2, 7, 9, 11, 12, 33, 34}
	utils.AssertArrayDeepEqual(t, val, expected)

	val = []int{33, 33, 33, 12, 2, 11, 9, 34, 7,}
	HeapSort(val)
	expected = []int{2, 7, 9, 11, 12, 33, 33, 33, 34}
	utils.AssertArrayDeepEqual(t, val, expected)

	val = []int{1, 2, 3, 4, 5}
	HeapSort(val)
	expected = []int{1, 2, 3, 4, 5}
	utils.AssertArrayDeepEqual(t, val, expected)

	val = []int{100, 99, 88, 77, 66, 55}
	HeapSort(val)
	expected = []int{55, 66, 77, 88, 99, 100}
	utils.AssertArrayDeepEqual(t, val, expected)

}

package exercise_2_1_2

import (
	"testing"
	utils "common/testUtils"
)

func TestInsertionSort(t *testing.T) {
	InsertionSortDesc(nil)
	InsertionSortDesc([]int{})
	InsertionSortDesc([]int{23})

	actual := []int{23, 43, 67, 1, 45, 8}
	expected := []int{67, 45, 43, 23, 8, 1}
	InsertionSortDesc(actual)
	utils.AssertArrayDeepEqual(t, actual, expected)

	actual = []int{1, 2, 3, 4, 5, 6}
	expected = []int{6, 5, 4, 3, 2, 1}
	InsertionSortDesc(actual)
	utils.AssertArrayDeepEqual(t, actual, expected)

	actual = []int{6, 5, 4, 3, 2, 1}
	expected = []int{6, 5, 4, 3, 2, 1}
	InsertionSortDesc(actual)
	utils.AssertArrayDeepEqual(t, actual, expected)

	actual = []int{2, 3, 4, 4, 4, 5}
	expected = []int{5, 4, 4, 4, 3, 2}
	InsertionSortDesc(actual)
	utils.AssertArrayDeepEqual(t, actual, expected)
}

package exercise_2_1_4

import (
	"testing"
	utils "common/testUtils"
)

func TestBinaryArrayAddition(t *testing.T) {
	a, b := []int{1, 1, 1}, []int{1, 1, 1}
	actual := BinaryArrayAddition(a, b)
	expected := []int{0, 1, 1, 1}
	utils.AssertArrayDeepEqual(t, actual, expected)

	a, b = []int{1, 1, 1}, []int{1, 1, 0}
	actual = BinaryArrayAddition(a, b)
	expected = []int{0, 1, 0, 1}
	utils.AssertArrayDeepEqual(t, actual, expected)
}

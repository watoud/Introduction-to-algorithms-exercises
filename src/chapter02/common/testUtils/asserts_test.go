package testing

import (
	"testing"
)

func TestDeepArrayEquals(t *testing.T) {

	a, b := []int{1, 2, 3}, []int{1, 2, 3}
	AssertArrayDeepEqual(t, a, b)

	a, b = []int{1, 2, 3}, []int{1, 2, 4}
	AssertArrayDeepEqual(t, a, b)
}

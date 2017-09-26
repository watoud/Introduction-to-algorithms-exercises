package PriorityHeap

import (
	"testing"

	utils "common/testUtils"
)

func TestPriorityHeapExtract(t *testing.T) {
	h := PHeap{
		make([]int, 5),
		0,
	}
	h.Insert(4)
	h.Insert(1)
	h.Insert(7)
	h.Insert(5)
	h.Insert(2)
	utils.AssertArrayDeepEqual(t, h.Extract(), 1)
	utils.AssertArrayDeepEqual(t, h.Extract(), 2)
	utils.AssertArrayDeepEqual(t, h.Extract(), 4)
	utils.AssertArrayDeepEqual(t, h.Extract(), 5)
	utils.AssertArrayDeepEqual(t, h.Extract(), 7)

}

func TestPriorityHeapInsert(t *testing.T) {
	h := PHeap{
		make([]int, 5),
		0,
	}

	h.Insert(4)
	utils.AssertArrayDeepEqual(t, h.Elements[0], 4)
	h.Insert(1)
	utils.AssertArrayDeepEqual(t, h.Elements[0], 1)
	h.Insert(7)
	utils.AssertArrayDeepEqual(t, h.Elements[0], 1)
	h.Insert(5)
	utils.AssertArrayDeepEqual(t, h.Elements[0], 1)
	h.Insert(2)
	utils.AssertArrayDeepEqual(t, h.Elements[0], 1)
}

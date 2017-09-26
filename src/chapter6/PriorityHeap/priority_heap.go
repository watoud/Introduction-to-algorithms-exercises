package PriorityHeap

// PHeap 最小优先队列
type PHeap struct {
	Elements []int // 保存优先队列元素
	Size     int   // 优先队列大小
}

// Extract 提取最优元素同事保持优先队列特性
func (ph *PHeap) Extract() int {
	if ph.Size <= 0 {
		panic("the heap is empty")
	}
	ret := ph.Elements[0]
	ph.Size--
	ph.Elements[0] = ph.Elements[ph.Size]

	MinHeapfy(ph.Elements, ph.Size, 0)
	return ret
}

func MinHeapfy(vals []int, size int, index int) {
	if index < 0 || index >= size {
		return
	}

	left, right, min := index*2+1, index*2+2, index
	if left < size && vals[left] < vals[min] {
		min = left
	}

	if right < size && vals[right] < vals[min] {
		min = right
	}

	if min != index {
		tmp := vals[index]
		vals[index] = vals[min]
		vals[min] = tmp
		MinHeapfy(vals, size, min)
	}
}

// Insert 插入元素同事保持优先队列特性
func (h *PHeap) Insert(val int) {
	if h.Size+1 > len(h.Elements) {
		tmp := make([]int, len(h.Elements)*2)
		copy(tmp, h.Elements)
		h.Elements = tmp
	}
	h.Elements[h.Size] = val
	h.Size++

	Heapfy(h.Elements, h.Size, h.Size-1)
}

// Heapfy 保持堆的特性
func Heapfy(elems []int, size int, index int) {
	if index < 0 || index >= size {
		return
	}

	for parent := (index+1)/2 - 1; parent >= 0 && elems[index] < elems[parent]; {
		tmp := elems[index]
		elems[index] = elems[parent]
		elems[parent] = tmp

		index = parent
	}
}

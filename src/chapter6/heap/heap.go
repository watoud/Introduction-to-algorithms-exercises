package heap

// HeapSort 堆排序
func HeapSort(val []int) {
	buildMaxHeap(val, len(val))
	for i := len(val) - 1; i > 0; i-- {
		tmp := val[i]
		val[i] = val[0]
		val[0] = tmp
		maxHeapfy(val, i, 0)
	}
}

/*
 * BuildMaxHeap 把一个数组建成最大堆
 */
func buildMaxHeap(val []int, size int) {
	begin := size / 2

	for begin >= 0 {
		maxHeapfy(val, size, begin)
		begin--
	}
}

/*
 * MaxHeapfy 维护堆的性质
 * 输入约束： index的左右子树都是最大堆，只有index可能不符合最大堆性质
 * 输出： 整棵树为最大堆
 */
func maxHeapfy(val []int, size int, index int) {
	if index < 0 || index >= size {
		return
	}
	left, right, largest := index*2+1, index*2+2, index
	if left < size && val[left] > val[largest] {
		largest = left
	}

	if right < size && val[right] > val[largest] {
		largest = right
	}

	if largest != index {
		tmp := val[index]
		val[index] = val[largest]
		val[largest] = tmp
		maxHeapfy(val, size, largest)
	}
}

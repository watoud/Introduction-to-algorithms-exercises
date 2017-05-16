package exercise_2_1_2

/*
	非升序排序
 */
func InsertionSortDesc(dat []int) {
	if dat == nil || len(dat) == 0 {
		return
	}

	length := len(dat)

	for i := 1; i < length; i++ {
		val := dat[i]
		j := i - 1
		for ; j >= 0 && val > dat[j]; j-- {
			dat[j+1] = dat[j]
		}
		dat[j+1] = val
	}
}

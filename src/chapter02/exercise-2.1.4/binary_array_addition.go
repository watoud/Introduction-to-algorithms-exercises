package exercise_2_1_4

func BinaryArrayAddition(a []int, b []int) []int {
	if a == nil || b == nil || len(a) != len(b) {
		return nil
	}

	if len(a) == 0 {
		return []int{}
	}

	length, carry := len(a), 0
	ret := make([]int, length+1)

	for i := 0; i < length; i++ {
		ret[i] = a[i] + b[i] + carry
		if ret[i] >= 2 {
			carry = 1
			ret[i] -= 2
		} else {
			carry = 0
		}
	}
	if carry == 1 {
		ret[length] = 1
	}

	return ret[:]
}

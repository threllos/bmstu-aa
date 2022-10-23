package utils

// return min int
func MinFromSome(values ...int) int {
	if len(values) == 0 {
		return 0
	}

	res := values[0]

	for _, val := range values {
		if val < res {
			res = val
		}
	}

	return res
}

// return |x|
func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

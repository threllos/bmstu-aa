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

// return min int
func MaxFromSome(values ...int) int {
	if len(values) == 0 {
		return 0
	}

	res := values[0]

	for _, val := range values {
		if val > res {
			res = val
		}
	}

	return res
}

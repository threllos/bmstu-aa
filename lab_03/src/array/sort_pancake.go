package array

func (a Array) SortPancake() {
	for cur_n := a.Len(); cur_n > 1; cur_n-- {
		i := a.findMaxIndex(cur_n)

		if i != cur_n-1 {
			a.flip(i)
			a.flip(cur_n - 1)
		}
	}
}

func (a Array) findMaxIndex(len int) int {
	max_i := 0

	for i := 0; i < len; i++ {
		if a.Less(max_i, i) { max_i = i }
	}

	return max_i
}

func (a Array) flip(i int) {
	start := 0

	for start < i {
		a.Swap(start, i)
		start++
		i--
	}
}

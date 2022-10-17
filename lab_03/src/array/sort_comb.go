package array

func (a Array) SortComb() {
	gap := a.Len()
	flag := true

	for gap != 1 || flag {
		nextGap(&gap)
		flag = false

		for i := 0; i < a.Len()-gap; i++ {
			if a.Less(i+gap, i) {
				a.Swap(i+gap, i)
				flag = true
			}
		}
	}
}

func nextGap(gap *int) {
	*gap = (*gap * 10) / 13
	if *gap < 1 {
		*gap = 1
	}
}
package array

func (a Array) SortComb() {
	step := a.Len() - 1;

	for step >= 1 {
		for i := 0; i + step < a.Len(); i++ {
			if (a.Less(i + step, i)) { a.Swap(i, i + step) }
		}

		step = step * 10 / 13;
	}
}
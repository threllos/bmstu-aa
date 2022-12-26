package array

func (a Array) SortInsert() {
	for i := 1; i < a.Len(); i++ {
		key := a[i]
		j := i - 1

		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}

		a[j+1] = key
	}
}

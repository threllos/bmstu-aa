package utils

import "testing"

var testTable = []struct {
	title string
	in    []int
	min   int
	max   int
}{
	{
		title: "default",
		in:    []int{3, 6, 1, 18},
		min:   1,
		max:   18,
	},
	{
		title: "empty",
		in:    []int{},
		min:   0,
		max:   0,
	},
	{
		title: "alone num",
		in:    []int{555},
		min:   555,
		max:   555,
	},
	{
		title: "repeat num",
		in:    []int{7, 7, 7, 7, 7, 7},
		min:   7,
		max:   7,
	},
	{
		title: "with neg",
		in:    []int{3, 5, -2, 777, -5, 12},
		min:   -5,
		max:   777,
	},
	{
		title: "all negatives",
		in:    []int{-3, -6, -12, -2, -1, -8, -11, -8},
		min:   -12,
		max:   -1,
	},
}

func TestMaxFromSome(t *testing.T) {
	for _, test := range testTable {
		// Arrange

		// Act
		res := MaxFromSome(test.in...)

		// Assert
		if test.max != res {
			t.Errorf("Incorrect result.\ntitle: %v\nin: %v\nout: %v\nres: %v\n",
				test.title, test.in, test.max, res)
		} else {
			t.Logf("Test pass '%v'.\n", test.title)
		}
	}
}

func TestMinFromSome(t *testing.T) {
	for _, test := range testTable {
		// Arrange

		// Act
		res := MinFromSome(test.in...)

		// Assert
		if test.min != res {
			t.Errorf("Incorrect result.\ntitle: %v\nin: %v\nout: %v\nres: %v\n",
				test.title, test.in, test.min, res)
		} else {
			t.Logf("Test pass '%v'.\n", test.title)
		}
	}
}

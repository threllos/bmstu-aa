package utils

import "testing"

var testAbsTable = []struct {
	title string
	in    int
	out   int
}{
	{
		title: "zero",
		in:    0,
		out:   0,
	},
	{
		title: "negative",
		in:    -5,
		out:   5,
	},
	{
		title: "positive",
		in:    17,
		out:   17,
	},
}

var testMinSomeTable = []struct {
	title string
	in    []int
	out   int
}{
	{
		title: "default",
		in:    []int{3, 6, 1, 18},
		out:   1,
	},
	{
		title: "empty",
		in:    []int{},
		out:   0,
	},
	{
		title: "alone num",
		in:    []int{555},
		out:   555,
	},
	{
		title: "repeat num",
		in:    []int{7, 7, 7, 7, 7, 7},
		out:   7,
	},
	{
		title: "with neg",
		in:    []int{3, 5, -2, 777, -5, 12},
		out:   -5,
	},
	{
		title: "all negatives",
		in:    []int{-3, -6, -12, -2, -1, -8, -11, -8},
		out:   -12,
	},
}

func TestAbs(t *testing.T) {
	for _, test := range testAbsTable {
		// Arrange

		// Act
		res := Abs(test.in)

		// Assert
		if test.out != res {
			t.Errorf("Incorrect result.\ntitle: %v\nin: %v\nout: %v\nres: %v\n",
				test.title, test.in, test.out, res)
		} else {
			t.Logf("Test pass '%v'.\n", test.title)
		}
	}
}

func TestMinFromSome(t *testing.T) {
	for _, test := range testMinSomeTable {
		// Arrange

		// Act
		res := MinFromSome(test.in...)

		// Assert
		if test.out != res {
			t.Errorf("Incorrect result.\ntitle: %v\nin: %v\nout: %v\nres: %v\n",
				test.title, test.in, test.out, res)
		} else {
			t.Logf("Test pass '%v'.\n", test.title)
		}
	}
}

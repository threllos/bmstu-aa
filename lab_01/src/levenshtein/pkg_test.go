package levenshtein

import "testing"

var testLevenshteinTable = []struct {
	title string
	in    []string
	lOut  int
	dlOut int
}{
	{
		title: "equal strings",
		in:    []string{"aaasssddd", "aaasssddd"},
		lOut:  0,
		dlOut: 0,
	},
	{
		title: "str1 smaller",
		in:    []string{"aaasssdd", "aaasssddd"},
		lOut:  1,
		dlOut: 1,
	},
	{
		title: "str2 smaller",
		in:    []string{"aaasssddd", "aaasss"},
		lOut:  3,
		dlOut: 3,
	},
	{
		title: "str1 empty",
		in:    []string{"", "qwer"},
		lOut:  4,
		dlOut: 4,
	},
	{
		title: "str2 empty",
		in:    []string{"rewq", ""},
		lOut:  4,
		dlOut: 4,
	},
	{
		title: "mirror str",
		in:    []string{"qwerty", "ytrewq"},
		lOut:  6,
		dlOut: 5,
	},
	{
		title: "double str",
		in:    []string{"asd", "aassdd"},
		lOut:  3,
		dlOut: 3,
	},
	{
		title: "with space",
		in:    []string{"ag qwe", "a wqe"},
		lOut:  3,
		dlOut: 2,
	},
	{
		title: "damerau specific",
		in:    []string{"qwertyuiop", "wqreytiupo"},
		lOut:  6,
		dlOut: 5,
	},
}

func TestLIterative(t *testing.T) {
	for _, test := range testLevenshteinTable {
		// Arrange
		l := Make(test.in[0], test.in[1])

		// Act
		res := l.LIterative()

		// Assert
		if test.lOut != res {
			t.Errorf("Incorrect result.\ntitle: %v\nin: %v\nout: %v\nres: %v\n",
				test.title, test.in, test.lOut, res)
		} else {
			t.Logf("Test pass '%v'.\n", test.title)
		}
	}
}

func TestDLIterative(t *testing.T) {
	for _, test := range testLevenshteinTable {
		// Arrange
		l := Make(test.in[0], test.in[1])

		// Act
		res := l.DLIterative()

		// Assert
		if test.dlOut != res {
			t.Errorf("Incorrect result.\ntitle: %v\nin: %v\nout: %v\nres: %v\n",
				test.title, test.in, test.dlOut, res)
		} else {
			t.Logf("Test pass '%v'.\n", test.title)
		}
	}
}

func TestDLRecursive(t *testing.T) {
	for _, test := range testLevenshteinTable {
		// Arrange
		l := Make(test.in[0], test.in[1])

		// Act
		res := l.DLRecursive()

		// Assert
		if test.dlOut != res {
			t.Errorf("Incorrect result.\ntitle: %v\nin: %v\nout: %v\nres: %v\n",
				test.title, test.in, test.dlOut, res)
		} else {
			t.Logf("Test pass '%v'.\n", test.title)
		}
	}
}

func TestDLRecursiveCash(t *testing.T) {
	for _, test := range testLevenshteinTable {
		// Arrange
		l := Make(test.in[0], test.in[1])

		// Act
		res := l.DLRecursiveCash()

		// Assert
		if test.dlOut != res {
			t.Errorf("Incorrect result.\ntitle: %v\nin: %v\nout: %v\nres: %v\n",
				test.title, test.in, test.dlOut, res)
		} else {
			t.Logf("Test pass '%v'.\n", test.title)
		}
	}
}

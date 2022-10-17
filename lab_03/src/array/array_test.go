package array

import "testing"

var testSortTable = []struct {
	title string
	in Array
	out Array
} {
	{
		title: "default",
		in: Array{1, 3, 5, 2, 4},
		out: Array{1, 2, 3, 4, 5},
	},
	{
		title: "asc",
		in: Array{1, 2, 3, 4, 5},
		out: Array{1, 2, 3, 4, 5},
	},
	{
		title: "desc",
		in: Array{5, 4, 3, 2, 1},
		out: Array{1, 2, 3, 4, 5},
	},
	{
		title: "equals",
		in: Array{1, 1, 1, 1, 1},
		out: Array{1, 1, 1, 1, 1},
	},
	{
		title: "alone",
		in: Array{777},
		out: Array{777},
	},
	{
		title: "empty",
		in: Array{},
		out: Array{},
	},
	{
		title: "with neg nums",
		in: Array{-1, 2, -3, 4, -5},
		out: Array{-5, -3, -1, 2, 4},
	},
	{
		title: "all neg nums",
		in: Array{-1, -2, -3, -4, -5},
		out: Array{-5, -4, -3, -2, -1},
	},
}

var testCompareTable = []struct {
	title string
	in []Array
	out bool
} {
	{
		title: "equal",
		in: []Array{{1, 3, 5, 2, 4}, {1, 3, 5, 2, 4}},
		out: true,
	},
	{
		title: "not equal",
		in: []Array{{1, 3, 5, 2, 4}, {4, 2, 5, 3, 1}},
		out: false,
	},
	{
		title: "empty array",
		in: []Array{{}, {}},
		out: true,
	},
	{
		title: "alone item array",
		in: []Array{{777}, {777}},
		out: true,
	},
	{
		title: "different len array",
		in: []Array{{1}, {1, 2, 3}},
		out: false,
	},
}

func TestSortComb(t *testing.T) {
	for _, test := range testSortTable {
		// Arrange
		arr := test.in.Copy()

		// Act
		t.Logf("starting test '%v'\n", test.title)
		arr.SortComb()

		// Assert
		if !arr.Compare(test.out) {
			t.Errorf("Incorrect result.\ntitle: %v\nin: %v\nout: %v\nres: %v\n",
				test.title, test.in, test.out, arr)
		}
	}
}

func TestSortInsert(t *testing.T) {
	for _, test := range testSortTable {
		// Arrange
		arr := test.in.Copy()

		// Act
		t.Logf("starting test '%v'\n", test.title)
		arr.SortInsert()

		// Assert
		if !arr.Compare(test.out) {
			t.Errorf("Incorrect result.\ntitle: %v\nin: %v\nout: %v\nres: %v\n",
				test.title, test.in, test.out, arr)
		}
	}
}

func TestSortPancake(t *testing.T) {
	for _, test := range testSortTable {
		// Arrange
		arr := test.in.Copy()

		// Act
		t.Logf("starting test '%v'\n", test.title)
		arr.SortPancake()

		// Assert
		if !arr.Compare(test.out) {
			t.Errorf("Incorrect result.\ntitle: %v\nin: %v\nout: %v\nres: %v\n",
				test.title, test.in, test.out, arr)
		}
	}
}

func TestCompare(t *testing.T) {
	for _, test := range testCompareTable {
		// Arrange
		arr1 := test.in[0].Copy()
		arr2 := test.in[1].Copy()

		// Act
		t.Logf("starting test '%v'\n", test.title)
		res := arr1.Compare(arr2)

		// Assert
		if res != test.out {
			t.Errorf("Incorrect result.\ntitle: %v\nin: [%v], [%v]\nout: %v\nres: %v\n",
				test.title, test.in[0], test.in[1], test.out, res)
		}
	}
}
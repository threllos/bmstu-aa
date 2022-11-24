package matrix

import "testing"

var testMulTable = []struct {
	title string
	in    [2]Matrix
	err   bool
	out   Matrix
}{
	{
		title: "default",
		in: [2]Matrix{
			{
				rows: 2, columns: 2,
				data: [][]int{
					{5, 2},
					{2, 5},
				},
			},
			{
				rows: 2, columns: 2,
				data: [][]int{
					{2, 5},
					{5, 2},
				},
			},
		},
		err: false,
		out: Matrix{
			rows: 2, columns: 2,
			data: [][]int{
				{20, 29},
				{29, 20},
			},
		},
	},
	{
		title: "not even lens",
		in: [2]Matrix{
			{
				rows: 3, columns: 3,
				data: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			{
				rows: 3, columns: 3,
				data: [][]int{
					{9, 8, 7},
					{6, 5, 4},
					{3, 2, 1},
				},
			},
		},
		err: false,
		out: Matrix{
			rows: 3, columns: 3,
			data: [][]int{
				{30, 24, 18},
				{84, 69, 54},
				{138, 114, 90},
			},
		},
	},
	{
		title: "1x1",
		in: [2]Matrix{
			{
				rows: 1, columns: 1,
				data: [][]int{
					{2},
				},
			},
			{
				rows: 1, columns: 1,
				data: [][]int{
					{5},
				},
			},
		},
		err: false,
		out: Matrix{
			rows: 1, columns: 1,
			data: [][]int{
				{10},
			},
		},
	},
	{
		title: "first matrix have more rows",
		in: [2]Matrix{
			{
				rows: 2, columns: 1,
				data: [][]int{
					{2},
					{1},
				},
			},
			{
				rows: 1, columns: 2,
				data: [][]int{
					{1, 2},
				},
			},
		},
		err: false,
		out: Matrix{
			rows: 2, columns: 2,
			data: [][]int{
				{2, 4},
				{1, 2},
			},
		},
	},
	{
		title: "first matrix have more columns",
		in: [2]Matrix{
			{
				rows: 2, columns: 3,
				data: [][]int{
					{1, 2, 2},
					{1, 2, 2},
				},
			},
			{
				rows: 3, columns: 2,
				data: [][]int{
					{1, 2},
					{1, 2},
					{1, 2},
				},
			},
		},
		err: false,
		out: Matrix{
			rows: 2, columns: 2,
			data: [][]int{
				{5, 10},
				{5, 10},
			},
		},
	},
	{
		title: "with neg nums",
		in: [2]Matrix{
			{
				rows: 3, columns: 3,
				data: [][]int{
					{1, -2, 3},
					{1, 2, 3},
					{1, 2, 3},
				},
			},
			{
				rows: 3, columns: 3,
				data: [][]int{
					{-1, 2, 3},
					{1, 2, 3},
					{1, 2, 3},
				},
			},
		},
		err: false,
		out: Matrix{
			rows: 3, columns: 3,
			data: [][]int{
				{0, 4, 6},
				{4, 12, 18},
				{4, 12, 18},
			},
		},
	},
	{
		title: "err in size",
		in: [2]Matrix{
			{
				rows: 1, columns: 2,
				data: [][]int{
					{1, 2},
				},
			},
			{
				rows: 1, columns: 2,
				data: [][]int{
					{1, 2},
				},
			},
		},
		err: true,
		out: Matrix{},
	},
}

func TestMulClassic(t *testing.T) {
	for _, test := range testMulTable {
		// Arrange

		// Act
		r, err := test.in[0].MulClassic(test.in[1])

		// Assert
		if (err == nil) == test.err && r.Compare(test.out) {
			t.Errorf("Incorrect result.\ntitle: %v\nin1: %v\nin2: %v\nout: %v\nres: %v\n",
				test.title, test.in[0], test.in[1], test.out, r)
		} else {
			t.Logf("Test pass '%v'.\n", test.title)
		}
	}
}

func TestMulWinograd(t *testing.T) {
	for _, test := range testMulTable {
		// Arrange

		// Act
		r, err := test.in[0].MulWinograd(test.in[1])

		// Assert
		if (err == nil) == test.err && r.Compare(test.out) {
			t.Errorf("Incorrect result.\ntitle: %v\nin1: %v\nin2: %v\nout: %v\nres: %v\n",
				test.title, test.in[0], test.in[1], test.out, r)
		} else {
			t.Logf("Test pass '%v'.\n", test.title)
		}
	}
}

func TestMulWinogradOptimized(t *testing.T) {
	for _, test := range testMulTable {
		// Arrange

		// Act
		r, err := test.in[0].MulWinogradOptimized(test.in[1])

		// Assert
		if (err == nil) == test.err && r.Compare(test.out) {
			t.Errorf("Incorrect result.\ntitle: %v\nin1: %v\nin2: %v\nout: %v\nres: %v\n",
				test.title, test.in[0], test.in[1], test.out, r)
		} else {
			t.Logf("Test pass '%v'.\n", test.title)
		}
	}
}

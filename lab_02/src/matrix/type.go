package matrix

import (
	"fmt"
	"math/rand"
)

type Matrix struct {
	rows    int
	columns int
	data    [][]int
}

func Create(r, c int) (Matrix, bool) {
	if r <= 0 || c <= 0 {
		return Matrix{}, false
	}

	m := Matrix{rows: r, columns: c}
	m.data = make([][]int, r)

	for i := 0; i < r; i++ {
		m.data[i] = make([]int, c)
	}

	return m, true
}

func (m Matrix) String() string {
	s := ""

	for _, r := range m.data {
		for _, el := range r {
			s += fmt.Sprintf("%d ", el)
		}
		s += "\n"
	}

	return s[:len(s)-1]
}

func (m *Matrix) RandomFill(seed int64) {
	rand.Seed(seed)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			m.data[i][j] = rand.Int() % 101
		}
	}
}

func (m *Matrix) UserFIll() error {
	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns; j++ {
			_, err := fmt.Scanf("%d", &m.data[i][j])

			if err != nil {
				return err
			}
		}

		fmt.Scanf("\n") // Reset buffer
	}

	return nil
}

func (m1 Matrix) Compare(m2 Matrix) bool {
	if m1.rows != m2.rows || m1.columns != m2.columns {
		return false
	}

	for i := 0; i < m1.rows; i++ {
		for j := 0; j < m1.columns; j++ {
			if m1.data[i][j] != m2.data[i][j] {
				return false
			}
		}
	}

	return true
}

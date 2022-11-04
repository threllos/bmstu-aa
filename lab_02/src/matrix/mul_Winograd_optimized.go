package matrix

import "fmt"

func (m Matrix) rowFactorOptimized() []int {
	factor := make([]int, m.rows)

	for i := 0; i < m.rows; i++ {
		for j := 1; j < m.columns; j += 2 {
			factor[i] += m.data[i][j-1] * m.data[i][j]
		}
	}

	return factor
}

func (m Matrix) columnFactorOptimized() []int {
	factor := make([]int, m.columns)

	for i := 0; i < m.columns; i++ {
		for j := 1; j < m.rows; j += 2 {
			factor[i] += m.data[j-1][i] * m.data[j][i]
		}
	}

	return factor
}

// следует использовать оптимизации:
// заменить операцию x = x + k; на x += k |
// заменить умножение на 2 на побитовый сдвиг |
// предвычислять некоторые слагаемые для алгоритма.
func (m1 Matrix) MulWinogradOptimized(m2 Matrix) (Matrix, error) {
	if m1.columns != m2.rows {
		return Matrix{}, fmt.Errorf("matrix1 columns not equal matrix2 rows")
	}

	res, _ := Create(m1.rows, m2.columns)
	rf := m1.rowFactorOptimized()
	cf := m2.columnFactorOptimized()

	for i := 0; i < m1.rows; i++ {
		for j := 0; j < m2.columns; j++ {
			res.data[i][j] -= rf[i] + cf[j]

			for k := 1; k < m1.columns; k += 2 {
				res.data[i][j] += ((m1.data[i][k-1] + m2.data[k][j]) *
					(m1.data[i][k] + m2.data[k-1][j]))
			}
		}
	}

	if m1.columns&1 != 0 {
		for i := 0; i < m1.rows; i++ {
			for j := 0; j < m2.columns; j++ {
				res.data[i][j] += (m1.data[i][m1.columns-1] *
					m2.data[m1.columns-1][j])
			}
		}
	}

	return res, nil
}

package matrix

import "fmt"

func (m Matrix) rowFactor() []int {
	factor := make([]int, m.rows)

	for i := 0; i < m.rows; i++ {
		for j := 0; j < m.columns/2; j++ {
			factor[i] = factor[i] + m.data[i][j*2+1]*m.data[i][j*2]
		}
	}

	return factor
}

func (m Matrix) columnFactor() []int {
	factor := make([]int, m.columns)

	for i := 0; i < m.columns; i++ {
		for j := 0; j < m.rows/2; j++ {
			factor[i] = factor[i] + m.data[j*2+1][i]*m.data[j*2][i]
		}
	}

	return factor
}

func (m1 Matrix) MulWinograd(m2 Matrix) (Matrix, error) {
	if m1.columns != m2.rows {
		return Matrix{}, fmt.Errorf("matrix1 columns not equal matrix2 rows")
	}

	res, _ := Create(m1.rows, m2.columns)
	rf := m1.rowFactor()
	cf := m2.columnFactor()

	for i := 0; i < res.rows; i++ {
		for j := 0; j < res.columns; j++ {
			res.data[i][j] = -rf[i] - cf[j]

			for k := 0; k < m1.columns/2; k++ {
				res.data[i][j] = res.data[i][j] + (m1.data[i][k*2+1]+
					m2.data[k*2][j])*(m1.data[i][k*2]+m2.data[k*2+1][j])
			}
		}
	}

	if m1.columns%2 != 0 {
		for i := 0; i < res.rows; i++ {
			for j := 0; j < res.columns; j++ {
				res.data[i][j] = res.data[i][j] + (m1.data[i][m1.columns-1] *
					m2.data[m2.rows-1][j])
			}
		}
	}

	return res, nil
}

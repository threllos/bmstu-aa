package levenshtein

import "lab_01/utils"

func (l Levenshtein) DLRecursiveCash() int {
	if l.isEmpty() {
		return l.ifEmpty()
	}

	l1, l2 := l.lens()
	c := make([][]int, l1) // cache

	for i := range c {
		c[i] = make([]int, l2)

		for j := range c[i] {
			c[i][j] = -1
		}
	}

	return l.dlRecursiveCash(c)
}

func (l Levenshtein) dlRecursiveCash(c [][]int) int {
	if l.isEmpty() {
		return l.ifEmpty()
	}

	l1, l2 := l.lens()

	if c[l1-1][l2-1] != -1 {
		return c[l1-1][l2-1]
	}

	eq := l.isEqualRunes(l1-1, l2-1)

	res := utils.MinFromSome(
		l.cutRune(1, 0).dlRecursiveCash(c)+1,
		l.cutRune(0, 1).dlRecursiveCash(c)+1,
		l.cutRune(1, 1).dlRecursiveCash(c)+eq,
	)

	if l.dlFlag(l1, l2) {
		res = utils.MinFromSome(
			res,
			l.cutRune(2, 2).DLRecursive()+1,
		)
	}

	c[l1-1][l2-1] = res

	return res

}

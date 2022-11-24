package levenshtein

import "lab_01/utils"

func (l Levenshtein) DLRecursive() int {
	l1, l2 := l.lens()

	if l.isEmpty() {
		return utils.MaxFromSome(l1, l2)
	}

	eq := l.isEqualRunes(l1-1, l2-1)

	res := utils.MinFromSome(
		l.cutRune(1, 0).DLRecursive()+1,
		l.cutRune(0, 1).DLRecursive()+1,
		l.cutRune(1, 1).DLRecursive()+eq,
	)

	if l.dlFlag(l1, l2) {
		res = utils.MinFromSome(
			res,
			l.cutRune(2, 2).DLRecursive()+1,
		)
	}

	return res
}

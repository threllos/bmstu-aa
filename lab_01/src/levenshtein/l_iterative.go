package levenshtein

import "lab_01/utils"

func (l Levenshtein) LIterative() int {
	if l.isEmpty() {
		return l.ifEmpty()
	}

	l1, l2 := l.lens()
	cbuf := make([]int, l1+1) // cur  buffer
	pbuf := make([]int, l1+1) // prev buffer

	for i := range pbuf {
		pbuf[i] = i
	}

	for i := 1; i < l2+1; i++ {
		cbuf[0] = i

		for j := 1; j < l1+1; j++ {
			eq := l.isEqualRunes(j-1, i-1)

			res := utils.MinFromSome(
				cbuf[j-1]+1,
				pbuf[j]+1,
				pbuf[j-1]+eq,
			)

			cbuf[j] = res
		}

		copy(pbuf, cbuf)
	}

	return cbuf[l1]
}

package levenshtein

import "testing"

var data = [3][2]string{
	{
		"dog",
		"god",
	},
	{
		"about",
		"above",
	},
	{
		"abbanition",
		"abaptiston",
	},
}

func BenchmarkLIterativeSmall(b *testing.B) {
	l := Make(data[0][0], data[0][1])
	for i := 0; i < b.N; i++ {
		l.LIterative()
	}
}

func BenchmarkDLIterativeSmall(b *testing.B) {
	l := Make(data[0][0], data[0][1])
	for i := 0; i < b.N; i++ {
		l.DLIterative()
	}
}

func BenchmarkDLRecursiveSmall(b *testing.B) {
	l := Make(data[0][0], data[0][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursive()
	}
}

func BenchmarkDLRecursiveCashSmall(b *testing.B) {
	l := Make(data[0][0], data[0][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursiveCash()
	}
}

func BenchmarkLIterativeMedium(b *testing.B) {
	l := Make(data[1][0], data[1][1])
	for i := 0; i < b.N; i++ {
		l.LIterative()
	}
}

func BenchmarkDLIterativeMedium(b *testing.B) {
	l := Make(data[1][0], data[1][1])
	for i := 0; i < b.N; i++ {
		l.DLIterative()
	}
}

func BenchmarkDLRecursiveMedium(b *testing.B) {
	l := Make(data[1][0], data[1][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursive()
	}
}

func BenchmarkDLRecursiveCashMedium(b *testing.B) {
	l := Make(data[1][0], data[1][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursiveCash()
	}
}

func BenchmarkLIterativeBig(b *testing.B) {
	l := Make(data[2][0], data[2][1])
	for i := 0; i < b.N; i++ {
		l.LIterative()
	}
}

func BenchmarkDLIterativeBig(b *testing.B) {
	l := Make(data[2][0], data[2][1])
	for i := 0; i < b.N; i++ {
		l.DLIterative()
	}
}

func BenchmarkDLRecursiveBig(b *testing.B) {
	l := Make(data[2][0], data[2][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursive()
	}
}

func BenchmarkDLRecursiveCashBig(b *testing.B) {
	l := Make(data[2][0], data[2][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursiveCash()
	}
}

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

func BenchmarkLIterativeLen3(b *testing.B) {
	l := Make(data[0][0], data[0][1])
	for i := 0; i < b.N; i++ {
		l.LIterative()
	}
}

func BenchmarkDLIterativeLen3(b *testing.B) {
	l := Make(data[0][0], data[0][1])
	for i := 0; i < b.N; i++ {
		l.DLIterative()
	}
}

func BenchmarkDLRecursiveLen3(b *testing.B) {
	l := Make(data[0][0], data[0][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursive()
	}
}

func BenchmarkDLRecursiveCashLen3(b *testing.B) {
	l := Make(data[0][0], data[0][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursiveCash()
	}
}

func BenchmarkLIterativeLen5(b *testing.B) {
	l := Make(data[1][0], data[1][1])
	for i := 0; i < b.N; i++ {
		l.LIterative()
	}
}

func BenchmarkDLIterativeLen5(b *testing.B) {
	l := Make(data[1][0], data[1][1])
	for i := 0; i < b.N; i++ {
		l.DLIterative()
	}
}

func BenchmarkDLRecursiveLen5(b *testing.B) {
	l := Make(data[1][0], data[1][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursive()
	}
}

func BenchmarkDLRecursiveCashLen5(b *testing.B) {
	l := Make(data[1][0], data[1][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursiveCash()
	}
}

func BenchmarkLIterativeLen10(b *testing.B) {
	l := Make(data[2][0], data[2][1])
	for i := 0; i < b.N; i++ {
		l.LIterative()
	}
}

func BenchmarkDLIterativeLen10(b *testing.B) {
	l := Make(data[2][0], data[2][1])
	for i := 0; i < b.N; i++ {
		l.DLIterative()
	}
}

func BenchmarkDLRecursiveLen10(b *testing.B) {
	l := Make(data[2][0], data[2][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursive()
	}
}

func BenchmarkDLRecursiveCashLen10(b *testing.B) {
	l := Make(data[2][0], data[2][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursiveCash()
	}
}

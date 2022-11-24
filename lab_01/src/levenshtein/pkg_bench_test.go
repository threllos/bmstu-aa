package levenshtein

import "testing"

type dataStrings [2]string

var data = map[string]dataStrings{
	"len5": {
		"mmxJQ",
		"C5sa4",
	},
	"len10": {
		"7zLicS3ZAa",
		"mpUvWJR6kT",
	},
	"len20": {
		"JqoW8gPge3IzNch1aMCw",
		"AI4Wtq06M36h4fahNGGm",
	},
	"len50": {
		"x8PA1muiVqBiGqhXRmuR9Ky0862xRkPULNvjOq8zMr7fpG4l1z",
		"sxSLnP38JBApwW8T3E0HRYHRoUa9i9KOgdO1DqlDfkkSSBVF1I",
	},
	"len100": {
		"xjibLrqa4T3ZHU5v233qskP5Wse3CKQDa2nZfUq32gWh6947w81APMOvG7mRLPKt7RL3MoT7RnIgYyfrvjCBx6Yw1TvSbQEoUaUm",
		"ilW0XcB8Iqz4FK9L4CkkUcST1s17eJbQyUQ2olP6xT5ehRtDlDfoJ5SkQ7Osp32uaXrFx2eZlRwSYqvXHILxTmeaTMCojoOLbGhs",
	},
	"len200": {
		"tlTI0xOwiqySik6iOvSXVJHRkAvneOZvCvl8BvEcaUlGH9FFcychLrUXTOQMcRqDd0MP9hxrt5dK0Lzt2HPMakTEXzcWiF3hHPwjKYqPMlvkidZyksG3wEG4lspDHWUk0RMPs0ZRYQqKTOqKtYaJEkhdDo0FrXyFpizAg5a6yZDJBslakwmBFG4RStrgipTT7P4NXnbu",
		"6fOpy9KbZHPfWqJsICI7KWjxpg5lohtIQ0MBVahlxVgGBwi9qmzTc9Wt0XfbDgLmEmR2Jdp1eH7FwqaeeClbXxR9FdImvsxIzYfpp0CSuWxAZ9CJKlahDpM8Dcok7s2ovi4a0iKqD9RgTlJ8SrBn7TJsTFJwILIk69jk8GhD2CcgejvBWIzbXLJRGLDAIkBGP2Mzsj7u",
	},
}

// Левенштейн нерекурсивно
func BenchmarkLIterativeLen5(b *testing.B) {
	l := Make(data["len5"][0], data["len5"][1])
	for i := 0; i < b.N; i++ {
		l.LIterative()
	}
}

func BenchmarkLIterativeLen10(b *testing.B) {
	l := Make(data["len10"][0], data["len10"][1])
	for i := 0; i < b.N; i++ {
		l.LIterative()
	}
}

func BenchmarkLIterativeLen20(b *testing.B) {
	l := Make(data["len20"][0], data["len20"][1])
	for i := 0; i < b.N; i++ {
		l.LIterative()
	}
}

func BenchmarkLIterativeLen50(b *testing.B) {
	l := Make(data["len50"][0], data["len50"][1])
	for i := 0; i < b.N; i++ {
		l.LIterative()
	}
}

func BenchmarkLIterativeLen100(b *testing.B) {
	l := Make(data["len100"][0], data["len100"][1])
	for i := 0; i < b.N; i++ {
		l.LIterative()
	}
}

func BenchmarkLIterativeLen200(b *testing.B) {
	l := Make(data["len200"][0], data["len200"][1])
	for i := 0; i < b.N; i++ {
		l.LIterative()
	}
}

// Дамерау-Левенштейн нерекурсивно
func BenchmarkDLIterativeLen5(b *testing.B) {
	l := Make(data["len5"][0], data["len5"][1])
	for i := 0; i < b.N; i++ {
		l.DLIterative()
	}
}

func BenchmarkDLIterativeLen10(b *testing.B) {
	l := Make(data["len10"][0], data["len10"][1])
	for i := 0; i < b.N; i++ {
		l.DLIterative()
	}
}

func BenchmarkDLIterativeLen20(b *testing.B) {
	l := Make(data["len20"][0], data["len20"][1])
	for i := 0; i < b.N; i++ {
		l.DLIterative()
	}
}
func BenchmarkDLIterativeLen50(b *testing.B) {
	l := Make(data["len50"][0], data["len50"][1])
	for i := 0; i < b.N; i++ {
		l.DLIterative()
	}
}
func BenchmarkDLIterativeLen100(b *testing.B) {
	l := Make(data["len100"][0], data["len100"][1])
	for i := 0; i < b.N; i++ {
		l.DLIterative()
	}
}
func BenchmarkDLIterativeLen200(b *testing.B) {
	l := Make(data["len200"][0], data["len200"][1])
	for i := 0; i < b.N; i++ {
		l.DLIterative()
	}
}

// Дамерау-Левеншткйн рекурсивно
func BenchmarkDLRecursiveLen5(b *testing.B) {
	l := Make(data["len5"][0], data["len5"][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursive()
	}
}

func BenchmarkDLRecursiveLen10(b *testing.B) {
	l := Make(data["len10"][0], data["len10"][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursive()
	}
}

// Дамерау-Левенштейн рекурсивно с кешем
func BenchmarkDLRecursiveCashLen5(b *testing.B) {
	l := Make(data["len5"][0], data["len5"][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursiveCash()
	}
}

func BenchmarkDLRecursiveCashLen10(b *testing.B) {
	l := Make(data["len10"][0], data["len10"][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursiveCash()
	}
}

func BenchmarkDLRecursiveCashLen20(b *testing.B) {
	l := Make(data["len20"][0], data["len20"][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursiveCash()
	}
}

func BenchmarkDLRecursiveCashLen50(b *testing.B) {
	l := Make(data["len50"][0], data["len50"][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursiveCash()
	}
}
func BenchmarkDLRecursiveCashLen100(b *testing.B) {
	l := Make(data["len100"][0], data["len100"][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursiveCash()
	}
}
func BenchmarkDLRecursiveCashLen200(b *testing.B) {
	l := Make(data["len200"][0], data["len200"][1])
	for i := 0; i < b.N; i++ {
		l.DLRecursiveCash()
	}
}

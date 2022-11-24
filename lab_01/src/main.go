package main

import (
	"fmt"
	"lab_01/levenshtein"
)

func calcDistances(l levenshtein.Levenshtein) {
	values := [4]int{}
	values[0] = l.LIterative()
	values[1] = l.DLIterative()
	values[2] = l.DLRecursive()
	values[3] = l.DLRecursiveCash()

	titles := [4]string{
		"Левенштейн нерекурсивно:",
		"Дамерау-Левенштейн нерекурсивно:",
		"Дамерау-Левенштейн рекурсивно:",
		"Дамерау-Левенштейн рекурсивно с кешем:",
	}

	for i := range values {
		fmt.Println(titles[i], values[i])
	}
}

func main() {
	var (
		err        error
		str1, str2 string
	)

	fmt.Print("String 1: ")
	_, err = fmt.Scan(&str1)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("String 2: ")
	_, err = fmt.Scan(&str2)

	if err != nil {
		fmt.Println(err)
		return
	}

	l := levenshtein.Make(str1, str2)

	calcDistances(l)
}

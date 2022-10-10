package main

import (
	"fmt"
	"lab_03/array"
	"sort"
	"flag"
	"time"
)

const REPEATS int = 1000

var rFlag = flag.Bool("r", false, "Random filling array")
var raFlag = flag.Bool("ra", false, "ASC filling array")
var rdFlag = flag.Bool("rd", false, "DESC filling array")
var repeatsFlag = flag.Int("repeats", REPEATS, "Repeats sort")


func benchmark(arr array.Array) error {
	durations := [4]time.Duration{}
	start := time.Time{}
	repeats := *repeatsFlag

	var arr1, arr2, arr3, arr4 array.Array

	for i := 0; i < repeats; i++ {
		arr1 = arr.Copy()
		arr2 = arr.Copy()
		arr3 = arr.Copy()
		arr4 = arr.Copy()

		start = time.Now()
		sort.Sort(arr1)
		durations[0] += time.Since(start)

		start = time.Now()
		arr2.SortComb()
		durations[1] += time.Since(start)

		start = time.Now()
		arr3.SortInsert()
		durations[2] += time.Since(start)

		start = time.Now()
		arr4.SortPancake()
		durations[3] += time.Since(start)
	}

	if !array.Compare(arr1, arr2, arr3, arr4) {
		return fmt.Errorf("arrays not equal after sort")
	}

	fmt.Printf("\n=== Benchmark ===\n\n")

	if (arr.Len() < 10) {
		fmt.Printf("Entry  array: %v\n", arr)
		fmt.Printf("Sorted array: %v\n\n", arr1)
	}

	fmt.Printf("Random : %v\n", *rFlag)
	fmt.Printf("Repeats: %d\n", repeats)
	fmt.Printf("Length : %d\n\n", arr.Len())

	fmt.Printf("default sort: %.2fns\n", float64(durations[0].Nanoseconds()) / float64(repeats))
	fmt.Printf("comb    sort: %.2fns\n", float64(durations[1].Nanoseconds()) / float64(repeats))
	fmt.Printf("insert  sort: %.2fns\n", float64(durations[2].Nanoseconds()) / float64(repeats))
	fmt.Printf("pancake sort: %.2fns\n", float64(durations[3].Nanoseconds()) / float64(repeats))

	fmt.Printf("\n===    end    ===\n")

	return nil
}

func main() {
	var (
		arr array.Array
		err error
		len int
	)

	flag.Parse()

	fmt.Print("Array len: ")
	_, err = fmt.Scan(&len)

	if err == nil && len <= 0 { err = fmt.Errorf("len under 1") }

	if err != nil {
		fmt.Println(err)
		return
	}

	arr = array.Make(len)

	switch {
	case *rFlag:
		arr.RandomFill(time.Now().Unix())
	case *raFlag:
		arr.ASCFill()
	case *rdFlag:
		arr.DESCFill()
	default:
		err = arr.UserFill()
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	err = benchmark(arr)

	if err != nil {
		fmt.Println(err)
		return
	}
}

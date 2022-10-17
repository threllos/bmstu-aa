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
	var (
		durations [4]time.Duration
		arrs [4]array.Array
		repeats = *repeatsFlag
	)

	for i := 0; i < repeats; i++ {
		arrs[0] = arr.Copy()
		arrs[1] = arr.Copy()
		arrs[2] = arr.Copy()
		arrs[3] = arr.Copy()

		start := time.Now()
		sort.Sort(arrs[0])
		durations[0] += time.Since(start)

		start = time.Now()
		arrs[1].SortComb()
		durations[1] += time.Since(start)

		start = time.Now()
		arrs[2].SortInsert()
		durations[2] += time.Since(start)

		start = time.Now()
		arrs[3].SortPancake()
		durations[3] += time.Since(start)
	}

	if !array.Compare(arrs[0], arrs[1], arrs[2], arrs[3]) {
		fmt.Println(arrs[0])
		fmt.Println(arrs[1])
		fmt.Println(arrs[2])
		fmt.Println(arrs[3])
		return fmt.Errorf("arrays not equal after sort")
	}

	fmt.Printf("\n=== Benchmark ===\n\n")

	if (arr.Len() <= 10) {
		fmt.Printf("Entry  array: %v\n", arr)
		fmt.Printf("Sorted array: %v\n\n", arrs[0])
	}

	fmt.Printf("Random : %v\n", *rFlag)
	fmt.Printf("Repeats: %d\n", repeats)
	fmt.Printf("Length : %d\n\n", arr.Len())

	repeatsI64 := int64(repeats);
	fmt.Printf("default sort: %dns\n", durations[0].Nanoseconds() / repeatsI64)
	fmt.Printf("comb    sort: %dns\n", durations[1].Nanoseconds() / repeatsI64)
	fmt.Printf("insert  sort: %dns\n", durations[2].Nanoseconds() / repeatsI64)
	fmt.Printf("pancake sort: %dns\n", durations[3].Nanoseconds() / repeatsI64)

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

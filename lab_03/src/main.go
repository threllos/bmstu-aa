package main

import (
	"flag"
	"fmt"
	"lab_03/array"
	"lab_03/benchmark"
	"time"
)

const REPEATS int = 1000

var rFlag = flag.Bool("r", false, "Random filling array")
var raFlag = flag.Bool("ra", false, "ASC filling array")
var rdFlag = flag.Bool("rd", false, "DESC filling array")
var repeatsFlag = flag.Int("repeats", REPEATS, "Repeats sort")

func main() {
	var (
		arr array.Array
		err error
		len int
	)

	flag.Parse()

	fmt.Print("Array len: ")
	_, err = fmt.Scan(&len)

	if err == nil && len <= 0 {
		err = fmt.Errorf("len under 1")
	}

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

	benchmark.Benchmark(arr, *repeatsFlag)
}

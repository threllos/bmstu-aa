package main

import (
	"flag"
	"fmt"
	"lab_02/benchmark"
	"lab_02/matrix"
	"log"
	"time"
)

const REPEATS int = 1000

var rFlag = flag.Bool("r", false, "Random filling array")
var repeatsFlag = flag.Int("repeats", REPEATS, "Repeats sort")

func main() {
	flag.Parse()
	var matrices [2]matrix.Matrix

	for i := 0; i < 2; i++ {
		fmt.Printf("Input matrix %d\n", i+1)

		var r, c int
		fmt.Print("rows, colums: ")
		_, err := fmt.Scanf("%d %d\n", &r, &c)
		if err != nil {
			log.Fatalln(err)
		}

		var ok bool
		matrices[i], ok = matrix.Create(r, c)
		if !ok {
			err = fmt.Errorf("cannot create matrix %d", i+1)
			log.Fatalln(err)
		}

		switch {
		case *rFlag:
			matrices[i].RandomFill(time.Now().Unix())
		default:
			fmt.Println("matrix:")
			err = matrices[i].UserFIll()

			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	// fmt.Println("m1:")
	// fmt.Println(matrices[0])
	// fmt.Println("m2:")
	// fmt.Println(matrices[1])
	// fmt.Println("res:")
	// res, err := matrices[0].MulClassic(matrices[1])
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(res)

	// fmt.Println("res 2:")
	// res, err = matrices[0].MulWinograd(matrices[1])
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(res)

	// fmt.Println("res 3:")
	// res, err = matrices[0].MulWinogradOptimized(matrices[1])
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(res)
	benchmark.Benchmark(matrices, *repeatsFlag)
}

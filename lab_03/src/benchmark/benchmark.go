package benchmark

import (
	"fmt"
	"lab_03/array"
)

// #include <inttypes.h>
// #include <stdint.h>
// uint64_t tick(void)
// {
//     uint32_t high, low;
//     __asm__ __volatile__ (
//         "rdtsc\n"
//         "movl %%edx, %0\n"
//         "movl %%eax, %1\n"
//         : "=r" (high), "=r" (low)
//         :: "%rax", "%rbx", "%rcx", "%rdx"
//         );
//     uint64_t ticks = ((uint64_t)high << 32) | low;
//     return ticks;
// }
import "C"

func Benchmark(arr array.Array, repeats int) {
	var durations [3]uint64
	var tmp array.Array

	for i := 0; i < repeats; i++ {
		tmp = arr.Copy()
		start := C.tick()
		tmp.SortComb()
		durations[0] += uint64(C.tick() - start)

		tmp = arr.Copy()
		start = C.tick()
		tmp.SortInsert()
		durations[1] += uint64(C.tick() - start)

		tmp = arr.Copy()
		start = C.tick()
		tmp.SortPancake()
		durations[2] += uint64(C.tick() - start)
	}

	fmt.Printf("\n=== Benchmark ===\n\n")

	if arr.Len() <= 10 {
		fmt.Printf("Entry  array: %v\n", arr)
		fmt.Printf("Sorted array: %v\n\n", tmp)
	}

	fmt.Printf("Repeats: %d\n", repeats)
	fmt.Printf("Length : %d\n\n", arr.Len())

	fmt.Printf("comb    sort: %dns\n", durations[0]/uint64(repeats))
	fmt.Printf("insert  sort: %dns\n", durations[1]/uint64(repeats))
	fmt.Printf("pancake sort: %dns\n", durations[2]/uint64(repeats))

	fmt.Printf("\n===    end    ===\n")
}

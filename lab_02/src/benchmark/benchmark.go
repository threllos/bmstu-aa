package benchmark

import (
	"fmt"
	"lab_02/matrix"
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

const TICKS_PER_NS int = 2

func Benchmark(mats [2]matrix.Matrix, repeats int) {
	var durations [3]uint64

	for i := 0; i < repeats; i++ {
		start := C.tick()
		mats[0].MulClassic(mats[1])
		durations[0] += uint64(C.tick() - start)

		start = C.tick()
		mats[0].MulWinograd(mats[1])
		durations[1] += uint64(C.tick() - start)

		start = C.tick()
		mats[0].MulWinogradOptimized(mats[1])
		durations[2] += uint64(C.tick() - start)
	}

	fmt.Printf("\n=== Benchmark ===\n\n")

	fmt.Printf("Repeats: %d\n\n", repeats)

	fmt.Printf("classic       mult: %dns\n", durations[0]/uint64(repeats*TICKS_PER_NS))
	fmt.Printf("winograd      mult: %dns\n", durations[1]/uint64(repeats*TICKS_PER_NS))
	fmt.Printf("winograd opt. mult: %dns\n", durations[2]/uint64(repeats*TICKS_PER_NS))

	fmt.Printf("\n===    end    ===\n")
}

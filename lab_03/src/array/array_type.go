package array

import (
	"math/rand"
	"fmt"
)

type Array []int


// Print interface
func (a Array) String() string {
	s := ""

	for i, val := range(a) {
		if i > 0 { s += " " }

		s += fmt.Sprint(val)
	}

	return s
}

// Sort interface
func (a Array) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
// Sort interface
func (a Array) Less(i, j int) bool { return a[i] < a[j] }
// Sort interface
func (a Array) Len() int { return len(a) }

// Copy array
func (a Array) Copy() Array {
	cp := Make(len(a))

	for i, val := range(a) { cp[i] = val }

	return cp
}


/*
Compare 2 arrays.
true  - equal,
false - not equal
*/
func (a Array) Compare(arr Array) bool {
	if a.Len() != arr.Len() { return false }

	for i := 0; i < a.Len(); i++ {
		if a[i] != arr[i] { return false }
	}

	return true
}

// Random filling array
func (arr Array) RandomFill(seed int64) {
	rand.Seed(seed)

	for i := 0; i < arr.Len(); i++ { arr[i] = rand.Int() }
}

// ASC filling array
func (arr Array) ASCFill() {
	for i := 0; i < arr.Len(); i++ { arr[i] = i }
}

// DESC filling array
func (arr Array) DESCFill() {
	for i := 0; i < arr.Len(); i++ { arr[i] = arr.Len() - i }
}

// User filling array from console
func (arr Array) UserFill() error {
	for i := 0; i < arr.Len(); i++ {
		_, err := fmt.Scan(&arr[i])

		if err != nil { return err }
	}

	return nil
}


// Make array with 'len' length
func Make(len int) Array { return make(Array, len) }

/*
Compare several arrays with Compare method.
true  - equal,
false - not equal
*/
func Compare(arrs ...Array) bool {
	if len(arrs) <= 0 {
		return false
	} else if len(arrs) == 1 {
		return true
	} else {
		flag := true
		def := arrs[0]

		for i := 1; i < len(arrs) && flag; i++ {
			flag = def.Compare(arrs[i])
		}

		return flag
	}
}

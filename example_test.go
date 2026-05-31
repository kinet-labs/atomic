package atomic_test

import (
	"fmt"

	"github.com/kinet-labs/atomic"
)

func Example() {
	// Uint32 is a thin wrapper around the primitive uint32 type.
	var atom atomic.Uint32

	// The wrapper ensures that all operations are atomic.
	atom.Store(42)
	fmt.Println(atom.Inc())
	fmt.Println(atom.CompareAndSwap(43, 0))
	fmt.Println(atom.Load())

	// Output:
	// 43
	// true
	// 0
}

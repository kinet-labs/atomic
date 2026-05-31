package atomic

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
)

// Pointer is an atomic pointer of type *T.
type Pointer[T any] struct {
	_ nocmp // disallow non-atomic comparison
	p atomic.Pointer[T]
}

// NewPointer creates a new Pointer.
func NewPointer[T any](v *T) *Pointer[T] {
	var p Pointer[T]
	if v != nil {
		p.p.Store(v)
	}
	return &p
}

// Load atomically loads the wrapped value.
func (p *Pointer[T]) Load() *T {
	return p.p.Load()
}

// Store atomically stores the passed value.
func (p *Pointer[T]) Store(val *T) {
	p.p.Store(val)
}

// Swap atomically swaps the wrapped pointer and returns the old value.
func (p *Pointer[T]) Swap(val *T) (old *T) {
	return p.p.Swap(val)
}

// CompareAndSwap is an atomic compare-and-swap.
func (p *Pointer[T]) CompareAndSwap(old, new *T) (swapped bool) {
	return p.p.CompareAndSwap(old, new)
}

// String returns a human readable representation of a Pointer's underlying value.
func (p *Pointer[T]) String() string {
	return fmt.Sprint(p.Load())
}

// MarshalJSON encodes the wrapped pointer into JSON.
func (p *Pointer[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.Load())
}

// UnmarshalJSON decodes JSON into the wrapped pointer.
func (p *Pointer[T]) UnmarshalJSON(b []byte) error {
	var v T
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	p.Store(&v)
	return nil
}

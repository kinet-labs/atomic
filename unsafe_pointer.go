package atomic

import (
	"sync/atomic"
	"unsafe"
)

// UnsafePointer is an atomic wrapper around unsafe.Pointer.
type UnsafePointer struct {
	_ nocmp // disallow non-atomic comparison

	v unsafe.Pointer
}

// NewUnsafePointer creates a new UnsafePointer.
func NewUnsafePointer(val unsafe.Pointer) *UnsafePointer {
	return &UnsafePointer{v: val}
}

// Load atomically loads the wrapped value.
func (p *UnsafePointer) Load() unsafe.Pointer {
	return atomic.LoadPointer(&p.v)
}

// Store atomically stores the passed value.
func (p *UnsafePointer) Store(val unsafe.Pointer) {
	atomic.StorePointer(&p.v, val)
}

// Swap atomically swaps the wrapped unsafe.Pointer and returns the old value.
func (p *UnsafePointer) Swap(val unsafe.Pointer) (old unsafe.Pointer) {
	return atomic.SwapPointer(&p.v, val)
}

// CAS is an atomic compare-and-swap.
//
// Deprecated: Use CompareAndSwap
func (p *UnsafePointer) CAS(old, new unsafe.Pointer) (swapped bool) {
	return p.CompareAndSwap(old, new)
}

// CompareAndSwap is an atomic compare-and-swap.
func (p *UnsafePointer) CompareAndSwap(old, new unsafe.Pointer) (swapped bool) {
	return atomic.CompareAndSwapPointer(&p.v, old, new)
}

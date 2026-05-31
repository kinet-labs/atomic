package atomic

import (
	"encoding/json"
	"sync"
)

// Atomic is a generic thread-safe wrapper for any type.
// It uses a read-write mutex for safe concurrent access.
type Atomic[T any] struct {
	_ nocmp // disallow non-atomic comparison

	lock  sync.RWMutex
	value T
}

// NewAtomic creates a new Atomic with the given initial value.
func NewAtomic[T any](value T) *Atomic[T] {
	return &Atomic[T]{value: value}
}

// Get atomically returns the current value.
func (a *Atomic[T]) Get() T {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.value
}

// Set atomically sets the value.
func (a *Atomic[T]) Set(value T) {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value = value
}

// Swap atomically sets the value and returns the old value.
func (a *Atomic[T]) Swap(value T) T {
	a.lock.Lock()
	defer a.lock.Unlock()
	old := a.value
	a.value = value
	return old
}

// Load is an alias for Get for compatibility with sync/atomic naming.
func (a *Atomic[T]) Load() T {
	return a.Get()
}

// Store is an alias for Set for compatibility with sync/atomic naming.
func (a *Atomic[T]) Store(value T) {
	a.Set(value)
}

// MarshalJSON implements json.Marshaler.
func (a *Atomic[T]) MarshalJSON() ([]byte, error) {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return json.Marshal(a.value)
}

// UnmarshalJSON implements json.Unmarshaler.
func (a *Atomic[T]) UnmarshalJSON(b []byte) error {
	a.lock.Lock()
	defer a.lock.Unlock()
	return json.Unmarshal(b, &a.value)
}

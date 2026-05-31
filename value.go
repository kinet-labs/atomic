package atomic

import "sync/atomic"

// Value shadows the type of the same name from sync/atomic
// https://godoc.org/sync/atomic#Value
type Value struct {
	_ nocmp // disallow non-atomic comparison

	atomic.Value
}

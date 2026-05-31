package atomic

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestUnsafePointer(t *testing.T) {
	i := int64(42)
	j := int64(0)
	k := int64(1)

	tests := []struct {
		desc      string
		newAtomic func() *UnsafePointer
		initial   unsafe.Pointer
	}{
		{
			desc: "non-empty",
			newAtomic: func() *UnsafePointer {
				return NewUnsafePointer(unsafe.Pointer(&i))
			},
			initial: unsafe.Pointer(&i),
		},
		{
			desc: "nil",
			newAtomic: func() *UnsafePointer {
				var p UnsafePointer
				return &p
			},
			initial: unsafe.Pointer(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			t.Run("Load", func(t *testing.T) {
				atom := tt.newAtomic()
				require.Equal(t, tt.initial, atom.Load(), "Load should report nil.")
			})

			t.Run("Swap", func(t *testing.T) {
				atom := tt.newAtomic()
				require.Equal(t, tt.initial, atom.Swap(unsafe.Pointer(&k)), "Swap didn't return the old value.")
				require.Equal(t, unsafe.Pointer(&k), atom.Load(), "Swap didn't set the correct value.")
			})

			t.Run("CAS", func(t *testing.T) {
				atom := tt.newAtomic()
				require.True(t, atom.CAS(tt.initial, unsafe.Pointer(&j)), "CAS didn't report a swap.")
				require.Equal(t, unsafe.Pointer(&j), atom.Load(), "CAS didn't set the correct value.")
			})

			t.Run("Store", func(t *testing.T) {
				atom := tt.newAtomic()
				atom.Store(unsafe.Pointer(&i))
				require.Equal(t, unsafe.Pointer(&i), atom.Load(), "Store didn't set the correct value.")
			})
		})
	}
}

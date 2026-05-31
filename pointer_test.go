package atomic

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPointer(t *testing.T) {
	type foo struct {
		V int `json:"v"`
	}

	i := foo{42}
	j := foo{0}
	k := foo{1}

	tests := []struct {
		desc      string
		newAtomic func() *Pointer[foo]
		initial   *foo
	}{
		{
			desc: "New",
			newAtomic: func() *Pointer[foo] {
				return NewPointer(&i)
			},
			initial: &i,
		},
		{
			desc: "New/nil",
			newAtomic: func() *Pointer[foo] {
				return NewPointer[foo](nil)
			},
			initial: nil,
		},
		{
			desc: "zero value",
			newAtomic: func() *Pointer[foo] {
				var p Pointer[foo]
				return &p
			},
			initial: nil,
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
				require.Equal(t, tt.initial, atom.Swap(&k), "Swap didn't return the old value.")
				require.Equal(t, &k, atom.Load(), "Swap didn't set the correct value.")
			})

			t.Run("CAS", func(t *testing.T) {
				atom := tt.newAtomic()
				require.True(t, atom.CompareAndSwap(tt.initial, &j), "CAS didn't report a swap.")
				require.Equal(t, &j, atom.Load(), "CAS didn't set the correct value.")
			})

			t.Run("Store", func(t *testing.T) {
				atom := tt.newAtomic()
				atom.Store(&i)
				require.Equal(t, &i, atom.Load(), "Store didn't set the correct value.")
			})
			t.Run("String", func(t *testing.T) {
				atom := tt.newAtomic()
				require.Equal(t, fmt.Sprint(tt.initial), atom.String(), "String did not return the correct value.")
			})

			t.Run("MarshalJSON", func(t *testing.T) {
				atom := tt.newAtomic()
				marshaledPointer, err := json.Marshal(atom)
				require.NoError(t, err)
				marshaledRaw, err := json.Marshal(tt.initial)
				require.NoError(t, err)
				require.Equal(t, marshaledRaw, marshaledPointer, "MarshalJSON did not return the correct value.")
			})
		})
	}

	t.Run("UnmarshalJSON", func(t *testing.T) {
		var p Pointer[foo]

		require.NoError(t, json.Unmarshal([]byte(`{"v":1024}`), &p))
		require.Equal(t, 1024, p.Load().V, "UnmarshalJSON should have expected result")
	})

	t.Run("UnmarshalJSON error", func(t *testing.T) {
		var p Pointer[foo]

		require.Error(t, json.Unmarshal([]byte(`{"v":true}`), &p), "json.Unmarshal should return an error")
	})
}

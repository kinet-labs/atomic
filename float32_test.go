package atomic

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFloat32(t *testing.T) {
	atom := NewFloat32(4.2)

	require.Equal(t, float32(4.2), atom.Load(), "Load didn't work.")

	require.True(t, atom.CAS(4.2, 0.5), "CAS didn't report a swap.")
	require.Equal(t, float32(0.5), atom.Load(), "CAS didn't set the correct value.")
	require.False(t, atom.CAS(0.0, 1.5), "CAS reported a swap.")

	atom.Store(42.0)
	require.Equal(t, float32(42.0), atom.Load(), "Store didn't set the correct value.")
	require.Equal(t, float32(42.5), atom.Add(0.5), "Add didn't work.")
	require.Equal(t, float32(42.0), atom.Sub(0.5), "Sub didn't work.")

	require.Equal(t, float32(42.0), atom.Swap(45.0), "Swap didn't return the old value.")
	require.Equal(t, float32(45.0), atom.Load(), "Swap didn't set the correct value.")

	t.Run("JSON/Marshal", func(t *testing.T) {
		atom.Store(42.5)
		bytes, err := json.Marshal(atom)
		require.NoError(t, err, "json.Marshal errored unexpectedly.")
		require.Equal(t, []byte("42.5"), bytes, "json.Marshal encoded the wrong bytes.")
	})

	t.Run("JSON/Unmarshal", func(t *testing.T) {
		err := json.Unmarshal([]byte("40.5"), &atom)
		require.NoError(t, err, "json.Unmarshal errored unexpectedly.")
		require.Equal(t, float32(40.5), atom.Load(), "json.Unmarshal didn't set the correct value.")
	})

	t.Run("JSON/Unmarshal/Error", func(t *testing.T) {
		err := json.Unmarshal([]byte("\"40.5\""), &atom)
		require.Error(t, err, "json.Unmarshal didn't error as expected.")
		assertErrorJSONUnmarshalType(t, err,
			"json.Unmarshal failed with unexpected error %v, want UnmarshalTypeError.", err)
	})

	t.Run("String", func(t *testing.T) {
		assert.Equal(t, "42.5", NewFloat32(42.5).String(),
			"String() returned an unexpected value.")
	})
}

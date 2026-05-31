package atomic

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDuration(t *testing.T) {
	atom := NewDuration(5 * time.Minute)

	require.Equal(t, 5*time.Minute, atom.Load(), "Load didn't work.")
	require.Equal(t, 6*time.Minute, atom.Add(time.Minute), "Add didn't work.")
	require.Equal(t, 4*time.Minute, atom.Sub(2*time.Minute), "Sub didn't work.")

	require.True(t, atom.CAS(4*time.Minute, time.Minute), "CAS didn't report a swap.")
	require.Equal(t, time.Minute, atom.Load(), "CAS didn't set the correct value.")

	require.Equal(t, time.Minute, atom.Swap(2*time.Minute), "Swap didn't return the old value.")
	require.Equal(t, 2*time.Minute, atom.Load(), "Swap didn't set the correct value.")

	atom.Store(10 * time.Minute)
	require.Equal(t, 10*time.Minute, atom.Load(), "Store didn't set the correct value.")

	t.Run("JSON/Marshal", func(t *testing.T) {
		atom.Store(time.Second)
		bytes, err := json.Marshal(atom)
		require.NoError(t, err, "json.Marshal errored unexpectedly.")
		require.Equal(t, []byte("1000000000"), bytes, "json.Marshal encoded the wrong bytes.")
	})

	t.Run("JSON/Unmarshal", func(t *testing.T) {
		err := json.Unmarshal([]byte("1000000000"), &atom)
		require.NoError(t, err, "json.Unmarshal errored unexpectedly.")
		require.Equal(t, time.Second, atom.Load(), "json.Unmarshal didn't set the correct value.")
	})

	t.Run("JSON/Unmarshal/Error", func(t *testing.T) {
		err := json.Unmarshal([]byte("\"1000000000\""), &atom)
		require.Error(t, err, "json.Unmarshal didn't error as expected.")
		assertErrorJSONUnmarshalType(t, err,
			"json.Unmarshal failed with unexpected error %v, want UnmarshalTypeError.", err)
	})

	t.Run("String", func(t *testing.T) {
		assert.Equal(t, "42s", NewDuration(42*time.Second).String(),
			"String() returned an unexpected value.")
	})
}

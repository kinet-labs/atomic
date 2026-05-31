package atomic

import (
	"encoding/json"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAtomicGeneric(t *testing.T) {
	t.Run("bool", func(t *testing.T) {
		a := NewAtomic(false)
		require.False(t, a.Get())
		a.Set(true)
		require.True(t, a.Get())
		old := a.Swap(false)
		require.True(t, old)
		require.False(t, a.Get())
	})

	t.Run("int", func(t *testing.T) {
		a := NewAtomic(42)
		require.Equal(t, 42, a.Get())
		a.Set(100)
		require.Equal(t, 100, a.Get())
	})

	t.Run("string", func(t *testing.T) {
		a := NewAtomic("hello")
		require.Equal(t, "hello", a.Get())
		a.Set("world")
		require.Equal(t, "world", a.Get())
	})

	t.Run("struct", func(t *testing.T) {
		type Config struct {
			Name  string
			Value int
		}
		a := NewAtomic(Config{Name: "test", Value: 1})
		require.Equal(t, "test", a.Get().Name)
		a.Set(Config{Name: "updated", Value: 2})
		require.Equal(t, "updated", a.Get().Name)
	})

	t.Run("load_store", func(t *testing.T) {
		a := NewAtomic(0)
		a.Store(42)
		require.Equal(t, 42, a.Load())
	})
}

func TestAtomicGenericConcurrent(t *testing.T) {
	a := NewAtomic(0)
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			a.Set(v)
			_ = a.Get()
		}(i)
	}
	wg.Wait()
}

func TestAtomicGenericJSON(t *testing.T) {
	t.Run("marshal", func(t *testing.T) {
		a := NewAtomic(42)
		b, err := json.Marshal(a)
		require.NoError(t, err)
		require.Equal(t, "42", string(b))
	})

	t.Run("unmarshal", func(t *testing.T) {
		a := NewAtomic(0)
		err := json.Unmarshal([]byte("100"), a)
		require.NoError(t, err)
		require.Equal(t, 100, a.Get())
	})

	t.Run("struct", func(t *testing.T) {
		type Data struct {
			X int `json:"x"`
		}
		a := NewAtomic(Data{X: 1})
		b, err := json.Marshal(a)
		require.NoError(t, err)
		require.Equal(t, `{"x":1}`, string(b))

		a2 := NewAtomic(Data{})
		err = json.Unmarshal([]byte(`{"x":42}`), a2)
		require.NoError(t, err)
		require.Equal(t, 42, a2.Get().X)
	})
}

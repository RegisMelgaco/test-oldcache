package oldcache

import (
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLinkedList(t *testing.T) {
	t.SkipNow()

	t.Run("given size 2, create with 2 item", func(t *testing.T) {
		ll := NewCache[int](2)

		assert.Equal(t, "-1 - 0, -1 - 0", ll.String())
	})

	t.Run("given size 3, create with 3 item", func(t *testing.T) {
		ll := NewCache[int](3)

		assert.Equal(t, "-1 - 0, -1 - 0, -1 - 0", ll.String())
	})
}

func TestLinkedListPut(t *testing.T) {
	t.SkipNow()

	slog.SetLogLoggerLevel(slog.LevelDebug)

	t.Run("given with space, add item without removing", func(t *testing.T) {
		ll := NewCache[int](3)

		ll.Put(0, 2)
		ll.Put(1, 3)
		ll.Put(2, 4)

		assert.Equal(t, "2 - 4, 1 - 3, 0 - 2", ll.String())
	})

	t.Run("given full list, add item removing oldest", func(t *testing.T) {
		ll := NewCache[int](3)

		ll.Put(0, 2)
		ll.Put(1, 3)
		ll.Put(2, 4)
		ll.Put(3, 5)

		assert.Equal(t, "3 - 5, 2 - 4, 1 - 3", ll.String())
	})

	t.Run("given repeated key, update value", func(t *testing.T) {
		ll := NewCache[int](3)

		ll.Put(0, 2)
		ll.Put(1, 3)
		ll.Put(2, 4)
		ll.Put(1, 5)

		assert.Equal(t, "1 - 5, 2 - 4, 1 - 3", ll.String())
	})
}

func TestLinkedListGet(t *testing.T) {
	slog.SetLogLoggerLevel(slog.LevelDebug)

	t.Run("given get call, order should be updated", func(t *testing.T) {
		ll := NewCache[int](3)

		ll.Put(0, 2)
		ll.Put(1, 3)
		ll.Put(2, 4)
		v, ok := ll.Get(0)

		assert.NotNil(t, v)
		assert.True(t, ok)

		assert.Equal(t, "0 - 2, 2 - 4, 1 - 3", ll.String())
	})
}

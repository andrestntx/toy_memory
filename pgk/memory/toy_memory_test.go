package memory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToyMemory(t *testing.T) {
	m := toyMemory{memory: map[string][]Location{}}

	t.Run("saving many values for the same key", func(t *testing.T) {
		m.Put("123", Location{
			Lat: 12.23,
			Lng: 11.44,
		})

		m.Put("123", Location{
			Lat: 12.23,
			Lng: 11.44,
		})

		m.Put("123", Location{
			Lat: 12.23,
			Lng: 11.44,
		})

		m.Put("123", Location{
			Lat: 15.23,
			Lng: 16.44,
		})

		assert.NotNil(t, m.memory)
		assert.NotNil(t, m.memory["123"])
		assert.Equal(t, 1, len(m.memory))
		assert.Equal(t, 4, len(m.memory["123"]))
	})

	t.Run("saving many values for different key", func(t *testing.T) {
		m.Put("124", Location{
			Lat: 12.23,
			Lng: 11.44,
		})

		m.Put("125", Location{
			Lat: 12.23,
			Lng: 11.44,
		})

		m.Put("126", Location{
			Lat: 12.23,
			Lng: 11.44,
		})

		assert.NotNil(t, m.memory)
		assert.NotNil(t, m.memory["123"])
		assert.NotNil(t, m.memory["124"])
		assert.NotNil(t, m.memory["125"])
		assert.NotNil(t, m.memory["126"])

		assert.Equal(t, 4, len(m.memory))
		assert.Equal(t, 3, len(m.memory["123"]))
		assert.Equal(t, 1, len(m.memory["124"]))
		assert.Equal(t, 1, len(m.memory["125"]))
		assert.Equal(t, 1, len(m.memory["126"]))
	})

	t.Run("getting all history", func(t *testing.T) {
		history := m.Get("123", 0)

		assert.NotNil(t, history)
		assert.Equal(t, 4, len(history))
	})

	t.Run("getting history with max", func(t *testing.T) {
		history := m.Get("123", 2)

		assert.NotNil(t, history)
		assert.Equal(t, 2, len(history))
	})

	t.Run("deleting history", func(t *testing.T) {
		m.Delete("123")

		assert.NotNil(t, m.memory)
		assert.Nil(t, m.memory["123"])

		result := m.Get("123", 0)

		assert.Nil(t, result)
	})
}

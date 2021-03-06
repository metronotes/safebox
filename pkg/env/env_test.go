package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	t.Run("Fetch unset environment variable", func(t *testing.T) {
		_, err := Fetch("UNSET_VARIABLE")
		assert.Error(t, err)
	})

	t.Run("Fetch empty environment variable", func(t *testing.T) {
		err := os.Setenv("EMPTY_VARIABLE", "")
		assert.NoError(t, err)
		_, err = Fetch("EMPTY_VARIABLE")
		assert.Error(t, err)
	})

	t.Run("Fetch valid environment variable", func(t *testing.T) {
		expected := "test"

		err := os.Setenv("VALID_ENV", expected)
		assert.NoError(t, err)

		value, err := Fetch("VALID_ENV")
		assert.NoError(t, err)

		assert.Equal(t, expected, value)
	})
}

func TestFetchDefault(t *testing.T) {
	expected := "default"
	res := FetchDefault("INVALID_ENV", expected)

	assert.Equal(t, expected, res)
}

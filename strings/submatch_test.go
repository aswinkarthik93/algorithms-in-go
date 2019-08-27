package strings_test

import (
	"testing"

	"github.com/aswinkarthik/algorithms-in-go/strings"
	"github.com/stretchr/testify/assert"
)

func TestSubmatch(t *testing.T) {
	t.Run("should return true if the given string is a substring", func(t *testing.T) {
		source := "this is the source string"
		substr := " is the so"

		assert.True(t, strings.Submatch(source, substr))
	})

	t.Run("should return false if the given string is not a substring", func(t *testing.T) {
		source := "this is the source string"
		substr := " is the fo"

		assert.False(t, strings.Submatch(source, substr))
	})

	t.Run("should handle bigger strings", func(t *testing.T) {
		source := "this is the source string"
		substr := "this is a bigger string not in the source"

		assert.False(t, strings.Submatch(source, substr))
	})

	t.Run("should handle empty strings", func(t *testing.T) {
		source := "this is the source string"
		substr := ""

		assert.True(t, strings.Submatch(source, substr))
	})

}

package emoji

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLookupEmoji(t *testing.T) {
	c := require.New(t)

	result, err := LookupEmoji("🐷")
	c.Nil(err)
	c.Equal("1F437", result.Key)
	c.Contains(result.Descriptor, "Pig")
}

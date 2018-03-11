package manpage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPaths(t *testing.T) {
	c := require.New(t)

	paths := Paths()
	c.True(len(paths) > 0)

	for _, path := range paths {
		_, err := os.Stat(path)
		c.Nil(err)
	}
}

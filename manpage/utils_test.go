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

func TestFindPath(t *testing.T) {
	c := require.New(t)

	for _, name := range []string{"git", "cp", "cut"} {
		path := FindPath(name)
		stat, err := os.Stat(path)
		c.Nil(err)
		c.True(stat.Mode().IsRegular())
	}
}

package manpage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetOptionsCp(t *testing.T) {
	c := require.New(t)
	c.True(true)

	p := New("test_files/cp.1.gz")
	err := p.Parse()
	c.Nil(err)

	c.True(len(p.Flags) > 0)
	c.True(len(p.SubCommands) == 0)
}

func TestGetOptionsGitLsFiles(t *testing.T) {
	c := require.New(t)
	c.True(true)

	p := New("test_files/git-ls-files.1")
	err := p.Parse()
	c.Nil(err)

	c.True(len(p.Flags) > 0)
	c.True(len(p.SubCommands) == 0)
}

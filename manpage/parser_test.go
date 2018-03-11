package manpage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetOptionsCp(t *testing.T) {
	c := require.New(t)
	c.True(true)

	p := New("/usr/share/man/man1/cp.1")
	err := p.Parse()
	c.Nil(err)

	opts := p.GetOptions()
	c.True(len(opts) > 0)
}

func TestGetOptionsGitLsFiles(t *testing.T) {
	c := require.New(t)
	c.True(true)

	Debug = true
	p := New("/usr/local/share/man/man1/git-ls-files.1")
	err := p.Parse()
	c.Nil(err)

	opts := p.GetOptions()
	c.True(len(opts) > 0)
}

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

	opts := p.GetOptions()
	c.True(len(opts) > 0)
}

func TestGetOptionsGitLsFiles(t *testing.T) {
	c := require.New(t)
	c.True(true)

	p := New("test_files/git-ls-files.1")
	err := p.Parse()
	c.Nil(err)

	opts := p.GetOptions()
	c.True(len(opts) > 0)
}

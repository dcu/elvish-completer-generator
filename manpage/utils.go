package manpage

import (
	"bytes"
	"os/exec"
	"strings"
)

// Paths returns the paths to query man pages
func Paths() []string {
	out, err := exec.Command("man", "--path").CombinedOutput()
	if err != nil {
		return []string{}
	}

	out = bytes.TrimSuffix(out, []byte("\n"))
	return strings.Split(string(out), ":")
}

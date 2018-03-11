package manpage

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
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

// FindPath finds the man page path for the given man page name
func FindPath(name string) string {
	{
		// if name is a file then return it
		stat, err := os.Stat(name)
		if err == nil && !stat.IsDir() {
			return name
		}
	}

	var manPagePath string
	for _, path := range Paths() {
		err := filepath.Walk(path, func(root string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if strings.Split(info.Name(), ".")[0] == name {
				manPagePath = root
				return io.EOF
			}

			return nil
		})

		if err == io.EOF {
			break
		}
	}

	return manPagePath
}

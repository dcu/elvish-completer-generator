package manpage

import (
	"strings"
)

// Tag represents a tag in the document
type Tag struct {
	Name    string
	Content []string
}

// IsFlag returns true if the tag is a flag
func (t *Tag) IsFlag() bool {
	if len(t.Content) == 0 {
		return false
	}

	return strings.HasPrefix(t.Content[0], "-")
}

// ToFlag returns the flag and the description
func (t *Tag) ToFlag() (string, string) {
	return t.Content[0], strings.Join(t.Content[1:], " ")
}

package manpage

import (
	"regexp"
	"strings"
)

var (
	flagRx = regexp.MustCompile(`(--?[\w-]+)`)
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
func (t *Tag) ToFlag() map[string]string {
	result := map[string]string{}
	flags := strings.Split(t.Content[0], ",")
	content := strings.Join(t.Content[1:], " ")
	for _, flag := range flags {
		matches := flagRx.FindStringSubmatch(flag)
		if len(matches) > 1 {
			result[matches[1]] = strings.TrimSpace(content)
		}
	}

	return result
}

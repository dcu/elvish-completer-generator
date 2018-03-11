package manpage

import (
	"regexp"
	"strings"
)

var (
	flagRx       = regexp.MustCompile(`\A(--?[\w-]+)`)
	subCommandRx = regexp.MustCompile(`\A([a-z_-]+)\z`)
)

// Tag represents a tag in the document
type Tag struct {
	Name    string
	Content []string
}

// ToOptions returns the flag and the description
func (t *Tag) ToOptions() map[string]string {
	result := map[string]string{}
	opts := strings.Split(t.Content[0], ",")
	content := strings.Join(t.Content[1:], " ")
	for _, opt := range opts {
		if matches := flagRx.FindStringSubmatch(opt); len(matches) > 1 {
			result[matches[1]] = strings.TrimSpace(content)
		} else if matches := subCommandRx.FindStringSubmatch(opt); t.Name == "TP" && len(matches) > 1 {
			result[matches[1]] = strings.TrimSpace(content)
		}
	}

	return result
}

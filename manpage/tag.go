package manpage

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	flagRx       = regexp.MustCompile(`\A(--?[\w-]+)`)
	subCommandRx = regexp.MustCompile(`\A([a-z]+[a-z_-]*)\S*\z`)
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
		if Debug {
			fmt.Printf("looking for flags or subcommands: %s\n", opt)
		}

		if matches := flagRx.FindStringSubmatch(opt); len(matches) > 1 {
			result[matches[1]] = strings.TrimSpace(content)
			if Debug {
				fmt.Printf("found flag %s\n", matches[1])
			}
		} else if matches := subCommandRx.FindStringSubmatch(opt); t.couldBeSubCommand() && len(matches) > 1 {
			result[matches[1]] = strings.TrimSpace(content)
			if Debug {
				fmt.Printf("found sub command %s\n", matches[1])
			}
		}
	}

	return result
}

func (t *Tag) couldBeSubCommand() bool {
	return t.Name == "TP" || t.Name == "PP"
}

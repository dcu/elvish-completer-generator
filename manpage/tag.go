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

// ToSubCommands returns the sub command and the description
func (t *Tag) ToSubCommands() []*SubCommand {
	result := []*SubCommand{}
	if !t.couldBeSubCommand() {
		return result
	}

	opts := strings.Split(t.Content[0], ",")
	content := strings.Join(t.Content[1:], " ")
	for _, opt := range opts {
		if Debug {
			fmt.Printf("looking for subcommands: %s\n", opt)
		}

		if matches := subCommandRx.FindStringSubmatch(opt); len(matches) > 1 {
			if Debug {
				fmt.Printf("found sub command %s\n", matches[1])
			}

			subCommand := &SubCommand{
				Name:        matches[1],
				Description: strings.TrimSpace(content),
			}

			result = append(result, subCommand)
		}
	}

	return result
}

// ToFlags returns the flag and the description
func (t *Tag) ToFlags() []*Flag {
	result := []*Flag{}
	opts := strings.Split(t.Content[0], ",")
	content := strings.Join(t.Content[1:], " ")
	for _, opt := range opts {
		if Debug {
			fmt.Printf("looking for flags: %s\n", opt)
		}

		if matches := flagRx.FindStringSubmatch(opt); len(matches) > 1 {
			if Debug {
				fmt.Printf("found flag %s\n", matches[1])
			}

			flag := &Flag{
				Name:        matches[1],
				Description: strings.TrimSpace(content),
			}

			result = append(result, flag)
		}
	}

	return result
}

func (t *Tag) couldBeSubCommand() bool {
	return t.Name == "TP" || t.Name == "PP"
}

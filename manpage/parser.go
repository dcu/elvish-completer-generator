package manpage

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	sectionRx = regexp.MustCompile(`(?i)\A\.(PP|IT|SH)\s?(.*)\z`)
	trimRx    = regexp.MustCompile(`\\f\w|\.(\w{2})\s?`)
	flagRx    = regexp.MustCompile(`(?i)\.?FL\s`)
)

var (
	// Debug enables/disables debug mode
	Debug = false
)

// Parser is a parser for man pages
type Parser struct {
	path string
	tags []*Tag
}

// New creates a new Parser
func New(path string) *Parser {
	parser := &Parser{
		path: path,
		tags: []*Tag{},
	}

	return parser
}

// GetOptions gets the options from the man page
func (p *Parser) GetOptions() map[string]string {
	opts := map[string]string{}

	for _, tag := range p.tags {
		if !tag.IsFlag() {
			if Debug && len(tag.Content) > 0 {
				fmt.Printf("Skipped: %s %#v\n", tag.Name, tag.Content)
			}
			continue
		}

		flag, desc := tag.ToFlag()
		opts[flag] = desc
	}

	return opts
}

// Parse parses the man page
func (p *Parser) Parse() error {
	f, err := os.Open(p.path)
	if err != nil {
		return err
	}

	defer func() { _ = f.Close() }()

	scanner := bufio.NewScanner(f)

	var content []string
	var section string

	p.tags = []*Tag{}
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, `.\`) || line == "." {
			// skip comment
			continue
		}

		matches := sectionRx.FindStringSubmatch(line)
		if len(matches) <= 1 {
			content = append(content, cleanContent(line))
			continue
		}

		if section != "" {
			p.tags = append(p.tags, &Tag{Name: section, Content: content})
		}

		if len(matches) > 1 {
			section = matches[1]
			content = []string{}
			if matches[2] != "" {
				content = append(content, cleanContent(matches[2]))
			}
		}

	}

	return nil
}

func cleanContent(content string) string {
	content = flagRx.ReplaceAllString(content, "-") // unescape flags
	content = trimRx.ReplaceAllString(content, "")  // remove extra tags

	return strings.TrimRight(content, " ") // some lines have extra spaces at the end
}

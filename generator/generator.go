package generator

import (
	"os"
	"strconv"

	"github.com/alecthomas/template"
	"github.com/dcu/elvish-completer-generator/types"
)

// Generator defines a completer generator
type Generator struct {
	CommandName string
	Flags       []*types.Flag
	SubCommands []*types.SubCommand
}

// New creates a new instance of a generator
func New(commandName string, flags []*types.Flag, subCommands []*types.SubCommand) *Generator {
	return &Generator{
		CommandName: commandName,
		Flags:       flags,
		SubCommands: subCommands,
	}
}

// Render renders the completer
func (g *Generator) Render() error {
	tmpl, err := template.New("completer.elv").Funcs(template.FuncMap{
		"quote": quote,
	}).Parse(_templateContent)
	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, g)
}

func quote(value string) (string, error) {
	return strconv.Quote(value), nil
}

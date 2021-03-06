// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/dcu/elvish-completer-generator/generator"
	"github.com/dcu/elvish-completer-generator/manpage"
	"github.com/spf13/cobra"
)

var (
	_dontCompleteFiles       bool
	_dontCompleteSubCommands bool
)

// manCmd represents the man command
var manCmd = &cobra.Command{
	Use:   "man <command>",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("missing command name")
		}

		path := manpage.FindPath(args[0])
		page := manpage.New(path)

		err := page.Parse()
		if err != nil {
			return err
		}

		if manpage.Debug {
			for _, sc := range page.SubCommands {
				fmt.Printf("%s        %s\n", sc.Name, sc.Description)
			}

			for _, flag := range page.Flags {
				fmt.Printf("%s        %s\n", flag.Name, flag.Description)
			}
		}

		fileName := page.Name + "-completer.elv"
		f, err := os.Create(fileName)
		if err != nil {
			return err
		}
		defer func() { _ = f.Close() }()

		gen := generator.New(page.Name, page.Flags, page.SubCommands)
		gen.DontCompleteFiles = _dontCompleteFiles
		gen.DontCompleteSubCommands = _dontCompleteSubCommands
		err = gen.Render(f)
		if err != nil {
			return err
		}

		fmt.Printf("Completer written to %s.\n", fileName)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(manCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// manCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	manCmd.Flags().BoolVarP(&manpage.Debug, "debug", "d", false, "Enables debug mode")

	manCmd.Flags().BoolVarP(&_dontCompleteFiles, "dont-complete-files", "F", false, "Disable auto-completing file names")
	manCmd.Flags().BoolVarP(&_dontCompleteSubCommands, "dont-complete-subcommands", "s", false, "Disable auto-complete sub commands")
}

package cmd

import (
	"fmt"
	"os"

	"github.com/ief2i-florent/go-f22/src/fs"
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:                   "copy [source] [target]",
	Short:                 "Permet de syncroniser 2 répertoire de manière recursive",
	Args:                  cobra.ExactArgs(2),
	DisableFlagParsing:    true,
	DisableFlagsInUseLine: true,
	DisableAutoGenTag:     true,
	Run: func(cmd *cobra.Command, args []string) {
		origin := args[0]
		target := args[1]

		if err := fs.Copy(origin, target); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		os.Exit(0)
	},
}

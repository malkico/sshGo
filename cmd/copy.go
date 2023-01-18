package cmd

import (
	"fmt"
	"os"

	"github.com/ief2i-florent/go-f22/src/fs"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var RootCommand = &cobra.Command{
	Use:   "copy <source> <target>",
	Short: "Allows you to synchronize 2 directories recursively",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		origin := args[0]
		target := args[1]

		files, err := fs.Scan(origin)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if fs.PROGRESS {
			fs.PROGRESS_BAR = progressbar.Default(files)
		}

		if err := fs.Copy(origin, target); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		os.Exit(0)
	},
}

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(removeCmd)
}

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove one or more notes",
	Long:  "Remove notes from the workspace if they exist",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "a note must be provided\n")
			os.Exit(1)
		}

		for _, name := range args {
			filename := fmt.Sprintf("%s.%s", name, "md")
			file := filepath.Join(workdir, filename)

			err := os.Remove(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "could not delete note %s: %v\n", name, err)
				os.Exit(1)
			}
			fmt.Fprintf(os.Stdout, "note %s removed\n", name)
		}
	},
}

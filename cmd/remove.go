package cmd

import (
	"fmt"
	"os"

	"github.com/hspazio/mint/storage"
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
		store, err := storage.NewStore()
		if err != nil {
			exit("could not initialize store", err)
		}
		if len(args) < 1 {
			exit("a note must be provided", nil)
		}

		for _, name := range args {
			if err := store.RemoveNote(name); err != nil {
				exit(fmt.Sprintf("could not delete note %s", name), err)
			}
			fmt.Fprintf(os.Stdout, "note %s removed\n", name)
		}
	},
}

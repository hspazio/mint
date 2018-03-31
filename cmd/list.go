package cmd

import (
	"fmt"

	"github.com/hspazio/mint/storage"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all notes",
	Long:  "List all notes available in the workspace",
	Run: func(cmd *cobra.Command, args []string) {
		store, err := storage.NewStore()
		if err != nil {
			exit("could not initialize store", err)
		}
		notes, err := store.Notes()
		if err != nil {
			exit("could not list notes", err)
		}
		printList(notes)
	},
}

func printList(notes []storage.Note) {
	for i, n := range notes {
		fmt.Printf("%d. %s\n", i, n.Name)
	}
}

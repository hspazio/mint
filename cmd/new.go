package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var name string

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the note")
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new note",
	Long:  "Create a new note",
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			fmt.Fprintf(os.Stderr, "a name must be provided\n")
			os.Exit(1)
		}
		if err := editNote(name); err != nil {
			fmt.Fprintf(os.Stderr, "could not start editor: %v\n", err)
			os.Exit(1)
		}
	},
}

func editNote(name string) error {
	filename := fmt.Sprintf("%s.%s", name, "md")
	edit := exec.Command(os.Getenv("EDITOR"), filepath.Join(workdir, filename))
	edit.Stdin = os.Stdin
	edit.Stdout = os.Stdout
	edit.Stderr = os.Stderr

	if err := edit.Start(); err != nil {
		return err
	}
	if err := edit.Wait(); err != nil {
		return err
	}
	return nil
}

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
	rootCmd.AddCommand(openCmd)
}

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Open a note for edit",
	Long:  "Open a note for edit, create a new one if it doesn't exist",
	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]
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

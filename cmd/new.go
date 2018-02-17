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
		if name != "" {
			name = fmt.Sprintf("%s.%s", name, "md")
			edit := exec.Command(os.Getenv("EDITOR"), filepath.Join(workdir, name))
			edit.Stdin = os.Stdin
			edit.Stdout = os.Stdout
			edit.Stderr = os.Stderr

			if err := edit.Start(); err != nil {
				fmt.Fprintf(os.Stderr, "could not start editor: %v", err)
				os.Exit(1)
			}

			if err := edit.Wait(); err != nil {
				fmt.Fprintf(os.Stderr, "could not wait for editor to finish: %v", err)
				os.Exit(1)
			}
		}
	},
}

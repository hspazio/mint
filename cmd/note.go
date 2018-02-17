package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var name string

func init() {
	rootCmd.AddCommand(noteCmd)
}

var noteCmd = &cobra.Command{
	Use:   "note",
	Short: "Create or edit a note",
	Long:  "Create or edit a note, create a new one if it doesn't exist",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "a name must be provided\n")
			os.Exit(1)
		}

		name := args[0]
		filename := fmt.Sprintf("%s.%s", name, "md")
		file := filepath.Join(workdir, filename)

		stdinInfo, err := os.Stdin.Stat()
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not read stats for Stdin: %v\n", err)
			os.Exit(1)
		}
		if stdinInfo.Size() > 0 {
			content, err := ioutil.ReadAll(os.Stdin)
			if err != nil {
				fmt.Fprintf(os.Stderr, "could not read from Stdin: %v\n", err)
				os.Exit(1)
			}
			err = ioutil.WriteFile(file, content, os.ModePerm)
			if err != nil {
				fmt.Fprintf(os.Stderr, "could not save file: %v\n", err)
				os.Exit(1)
			}
		} else {
			if err := edit(file); err != nil {
				fmt.Fprintf(os.Stderr, "could not start editor: %v\n", err)
				os.Exit(1)
			}
		}
	},
}

func edit(file string) error {
	edit := exec.Command(os.Getenv("EDITOR"), file)
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

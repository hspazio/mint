package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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
		files, err := ioutil.ReadDir(workdir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "could not list directory: %v\n", err)
			os.Exit(1)
		}
		for i, f := range files {
			name := strings.TrimSuffix(f.Name(), ".md")
			fmt.Printf("%d. %s\n", i+1, name)
		}
	},
}

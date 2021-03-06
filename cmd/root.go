package cmd

import (
	"fmt"
	"os"

	"github.com/hspazio/mint/storage"
	"github.com/spf13/cobra"
)

var store *storage.Store

var rootCmd = &cobra.Command{
	Use:   "mint",
	Short: "Easily manage ideas",
	Long:  "Mint is a simple comman line tool to manage notes",
}

// Execute root command
func Execute() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(noteCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(removeCmd)
	rootCmd.Execute()
}

func exit(msg string, err error) {
	fmt.Fprintf(os.Stderr, "%s: %s\n", msg, err)
	os.Exit(1)
}

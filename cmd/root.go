package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

var workdir = initWorkdir()

var rootCmd = &cobra.Command{
	Use:   "mint",
	Short: "Easily manage ideas",
	Long:  "Mint is a simple comman line tool to manage notes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mint 0.1 - using path", workdir)
	},
}

// Execute root command
func Execute() {
	rootCmd.Execute()
}

func initWorkdir() string {
	rootdir := "."

	usr, err := user.Current()
	if err == nil {
		rootdir = usr.HomeDir
	}

	dir := filepath.Join(rootdir, ".mint")
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		fmt.Fprintf(os.Stderr, "could not create directory %s: %v\n", dir, err)
		os.Exit(1)
	}
	return dir
}

package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

var workdir string

func init() {
	// TODO: remove or use config file
	rootCmd.PersistentFlags().StringVarP(&workdir, "workdir", "d", initWorkdir(), "Working directory")
}

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
		fmt.Fprintf(os.Stderr, "could not create directory %s: %v", dir, err)
		os.Exit(1)
	}
	return dir
}

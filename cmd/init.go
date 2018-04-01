package cmd

import (
	"fmt"

	"github.com/hspazio/mint/configurations"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize default configurations",
	Long:  "initialize default configurations",
	Run: func(cmd *cobra.Command, args []string) {
		if err := configurations.Init(); err != nil {
			exit("could not initialize configurations", err)
		}
		fmt.Println("done!")
	},
}

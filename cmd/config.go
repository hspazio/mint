package cmd

import (
	"fmt"
	"os"

	"github.com/hspazio/mint/configurations"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Read/write configurations",
	Long:  "Read/write configurations",
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			conf, err := configurations.GetAll()
			if err != nil {
				fmt.Println("no configurations found, run 'mint init' command")
				os.Exit(1)
			}
			fmt.Println(conf)
		case 2:
			if err := configurations.Set(args[0], args[1]); err != nil {
				exit("%v", err)
			}
		default:
			fmt.Println("mint config - prints all configurations")
			fmt.Println("mint config dir /path/to/dir - sets a conf value")
		}
	},
}

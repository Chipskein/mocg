package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill mocg server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Should kill mocg server")
	},
}

func init() {

	rootCmd.AddCommand(killCmd)
}

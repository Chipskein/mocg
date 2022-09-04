package cmd

import (
	"fmt"
	"os"

	"github.com/chipskein/mocg/internals/ui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mocg [command]",
	Short: "Music on console with Go",
	Long:  `Music on console with Go. Terminal music player inspired in MOC player`,
	Run: func(cmd *cobra.Command, args []string) {
		//verify if mocg subprocess is in background
		//if in restore
		//else start
		ui.StartUI()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

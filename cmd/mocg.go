package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/chipskein/mocg/internals/ui"
	"github.com/spf13/cobra"
)

var mocgCmd = &cobra.Command{
	Use:   "mocg [command]",
	Short: "Music on console with Go",
	Long:  `Music on console with Go. Terminal music player inspired in MOC player`,
	Run: func(cmd *cobra.Command, args []string) {
		dirname, err := os.UserHomeDir()
		if err != nil {
			log.Fatal(err)
		}
		var default_dir = fmt.Sprintf("%s/Music", dirname)
		ui.StartUI("Should just fail", default_dir, true)
	},
}

func Execute() {
	if err := mocgCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

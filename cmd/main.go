package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Init() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Init MOCG server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Init MOCG server")
		},
	}
}

func Kill() *cobra.Command {
	return &cobra.Command{
		Use:   "kill",
		Short: "Kill MOCG server",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Kill MOCG server")
		},
	}
}

func main() {
	rootCmd := cobra.Command{}
	rootCmd.AddCommand(Init(), Kill())
	rootCmd.Execute()
}

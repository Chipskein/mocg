package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Init() *cobra.Command {
	return &cobra.Command{
		Use:   "hello [name]",
		Short: "retorna Olá + name passado",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Olá %s\n", args[0])
		},
	}
}
func Restore() *cobra.Command {
	return &cobra.Command{
		Use:   "hello [name]",
		Short: "retorna Olá + name passado",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Olá %s\n", args[0])
		},
	}
}
func Kill() *cobra.Command {
	return &cobra.Command{
		Use:   "hello [name]",
		Short: "retorna Olá + name passado",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Olá %s\n", args[0])
		},
	}
}
func main() {
	rootCmd := cobra.Command{}
	rootCmd.AddCommand(Init(), Restore(), Kill())
	rootCmd.Execute()
}

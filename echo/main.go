package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "echo [string...]",
		Short: "Prints the given string to the console",
		Long:  "Prints the given string to the console",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("")
				return
			}
			fmt.Println(strings.Join(args[0:], " "))
		},
	}

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

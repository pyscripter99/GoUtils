package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "myapp",
		Short: "A simple command-line application",
		Long:  "A simple command-line application created using Cobra",
		Run: func(cmd *cobra.Command, args []string) {
			// Default action if no subcommand is specified
			fmt.Println("Hello, World!")
		},
	}

	// Add subcommand
	rootCmd.AddCommand(&cobra.Command{
		Use:   "greet [name]",
		Short: "Greet someone",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Hello, %s!\n", args[0])
		},
	})

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

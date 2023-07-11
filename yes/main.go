package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "yes [string]",
		Short: "Repeatedly output a line with specified string, or 'y'",
		Long:  "Repeatedly output a line with specified string, or 'y'",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				for {
					fmt.Println("y")
				}
			} else {
				for {
					fmt.Println(args[0])
				}
			}
		},
	}

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

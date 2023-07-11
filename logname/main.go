package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "logname",
		Short: "Print the user's login name",
		Long:  "Print the user's login name.",
		Run: func(cmd *cobra.Command, args []string) {
			// Get the current user username (cross-platform)
			user, err := user.Current()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Print the username
			fmt.Println(user.Username)
		},
	}

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

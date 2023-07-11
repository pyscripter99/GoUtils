package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "whoami",
		Short: "Print the current user",
		Long:  "Print the current user",
		Run: func(cmd *cobra.Command, args []string) {
			// Get the current user
			u, err := user.Current()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Print(u.Username)

			// Print hostname if flag is set
			printHostname, err := cmd.Flags().GetBool("hostname")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if printHostname {
				hostname, err := os.Hostname()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Printf("@%s\n", hostname)
			} else {
				fmt.Println() // Print newline because we didn't print hostname
			}
		},
	}

	// Add flags
	rootCmd.Flags().BoolP("hostname", "n", false, "Prints hostname and username, format: 'username@hostname'")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

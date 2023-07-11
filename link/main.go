package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "link <target> <linkname>",
		Short: "Creates a hard link to a file",
		Long:  "Creates a hard link to a file. If the linkname exists, an error is returned.",
		Run: func(cmd *cobra.Command, args []string) {
			// Check if all arguments were passed
			if len(args) <= 1 {
				fmt.Println("link: missing operand")
				os.Exit(1)
			} else if len(args) > 2 {
				fmt.Println("link: extra operand")
				os.Exit(1)
			}

			// Check if the symbolic flag was passed
			if symbolic, _ := cmd.Flags().GetBool("symbolic"); symbolic {
				// Create the symbolic link
				if err := os.Symlink(args[0], args[1]); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				return
			}

			// Create the hard link
			if err := os.Link(args[0], args[1]); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		},
	}

	// Add flags
	rootCmd.Flags().BoolP("symbolic", "s", false, "make symbolic links instead of hard links")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

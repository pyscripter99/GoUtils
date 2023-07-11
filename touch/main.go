package main

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "touch <file>",
		Short: "Updates the modification and access times of a file",
		Long:  "Updates the modification and access times of a file. If the file does not exist, it is created with default \npermissions. Uses the current time if not specified.",
		Run: func(cmd *cobra.Command, args []string) {
			// Check if no arguments were passed
			if len(args) == 0 {
				fmt.Println("touch: missing file operand")
				os.Exit(1)
			}

			// Check if the file exists
			if _, err := os.Stat(args[0]); os.IsNotExist(err) {
				// Check if the no-create flag was passed
				if noCreate, _ := cmd.Flags().GetBool("no-create"); noCreate {
					fmt.Println("touch: cannot touch 'file': No such file or directory")
					os.Exit(1)
				}

				// Create the file
				file, err := os.Create(args[0])
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				defer file.Close()
			}

			// Check if the date flag was passed
			if date, _ := cmd.Flags().GetString("date"); date != "" {
				// Parse the date
				t, err := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", date)
				if err != nil {
					fmt.Println(err)
					os.Exit(0)
				}

				// Update the file's modification and access times
				if err := os.Chtimes(args[0], t, t); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			} else if reference, _ := cmd.Flags().GetString("reference"); reference != "" {
				// Get reference file's dates
				info, err := os.Stat(reference)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if err := os.Chtimes(args[0], info.ModTime(), info.ModTime()); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			} else {
				// Update the file's modification and access times
				if err := os.Chtimes(args[0], time.Now(), time.Now()); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
			}
		},
	}

	// Add flags
	rootCmd.Flags().BoolP("no-create", "c", false, "Do not create any files.")
	rootCmd.Flags().StringP("date", "d", "", "Parse argument and use it instead of current time.")
	rootCmd.Flags().StringP("reference", "r", "", "Use this file's times instead of current time.")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

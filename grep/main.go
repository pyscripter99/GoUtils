package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func printLine(line string, regex string, showNoMatch bool) {
	reg, err := regexp.Compile(regex)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Print line if it matches
	if reg.MatchString(line) {
		// Find matches and highlight them
		matches := reg.FindAllStringIndex(line, -1)
		// Print line (match can be anywhere in line)
		for i, match := range matches {
			if i == 0 {
				color.New(color.Bold).Print(line[:match[0]])
			}
			color.New(color.Bold).Print(color.RedString(line[match[0]:match[1]]))
			if i == len(matches)-1 {
				color.New(color.Bold).Println(line[match[1]:])
			} else {
				color.New(color.Bold).Print(line[match[1]:matches[i+1][0]])
			}
		}
	} else if showNoMatch {
		// Print line if it doesn't match
		fmt.Println(line)
	}
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "grep <regex> [file...]",
		Short: "Finds lines matching a regular expression",
		Long:  "Finds lines matching a regular expression from standard input or a list of files",
		Run: func(cmd *cobra.Command, args []string) {
			// Check for regex
			if len(args) < 1 {
				fmt.Println("You must specify a regular expression")
				os.Exit(1)
			}

			// Check for files
			files := []string{}
			if len(args) > 1 {
				files = args[1:]
			}

			if len(files) == 0 {
				// Read from standard input
				stdin, err := io.ReadAll(os.Stdin)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				// remove trailing newline
				stdin = stdin[:len(stdin)-1]
				for _, line := range strings.Split(string(stdin), "\n") {
					// Get regex
					printLine(line, args[0], cmd.Flags().Changed("no-match"))
				}
			} else {
				// Read from files
				for _, file := range files {
					// Check if file exists or is a directory
					if _, err := os.Stat(file); os.IsNotExist(err) {
						fmt.Printf("File %s does not exist\n", file)
						os.Exit(1)
					} else if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					// Print filename
					color.New(color.FgCyan).Add(color.Bold).Println(file + ":")

					// Open file
					f, err := os.Open(file)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					// Read file
					lines, err := io.ReadAll(f)
					if err != nil {
						fmt.Println(err)
						os.Exit(1)
					}

					for _, line := range strings.Split(string(lines), "\n") {
						// Get regex
						printLine(line, args[0], cmd.Flags().Changed("no-match"))
					}
				}
			}
		},
	}

	// Add flags
	rootCmd.Flags().BoolP("no-match", "m", false, "Show lines that don't match")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

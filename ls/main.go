package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ls [directory...]",
	Short: "Lists files in specified directory",
	Long:  "Lists files in specified directory, by default current directory\nA filename will be green if it is executable, blue if it is a directory, and white otherwise.",
	Run: func(cmd *cobra.Command, args []string) {
		// Get Directory(s)
		dirs := []string{"."}
		if len(args) > 0 {
			dirs = args
		}

		// Loop through directories
		for _, dir := range dirs {
			// Open directory
			f, err := os.Open(dir)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Read directory
			names, err := f.Readdirnames(-1)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Print directory
			if len(dirs) > 1 {
				fmt.Println(dir)
			}
			for _, name := range names {
				toPrint := ""
				// Print in blue if directory
				if err != nil {
					toPrint = color.RedString("%s: %s", name, err)
					if len(dirs) > 1 {
						fmt.Println("  " + toPrint)
					} else {
						fmt.Println(toPrint)
					}
					continue
				}
				fi, err := os.Stat(filepath.Join(dir, name))
				if err != nil {
					toPrint = color.RedString("%s: %s", name, err)
					if len(dirs) > 1 {
						fmt.Println("  " + toPrint)
					} else {
						fmt.Println(toPrint)
					}
					continue
				}
				if fi.IsDir() {
					toPrint = color.BlueString("%s", name)
				} else {
					// if executable, print in green
					if fi.Mode()&0111 != 0 {
						toPrint = color.GreenString("%s", name)
					} else {
						toPrint = color.WhiteString("%s", name)
					}
				}
				if len(dirs) > 1 {
					fmt.Println("  " + toPrint)
				} else {
					fmt.Println(toPrint)
				}
			}
		}
	},
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

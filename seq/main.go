package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "seq [from] <to> [increment]",
		Short: "Print sequences of numbers",
		Long:  "Print sequences of numbers.",
		Run: func(cmd *cobra.Command, args []string) {
			// Check if to argument was passed
			if len(args) < 1 {
				fmt.Println("seq: missing operand")
				os.Exit(1)
			}

			// Check if all arguments are numbers
			for _, arg := range args {
				_, err := strconv.Atoi(arg)
				if err != nil {
					fmt.Println("seq: invalid number '" + arg + "'")
					os.Exit(1)
				}
			}

			from := 1
			to := 1
			increment := 1

			// if 1 argument is passed, assume it's the to argument
			if len(args) == 1 {
				to, _ = strconv.Atoi(args[0])
			}

			// if 2 arguments are passed, assume it's the from and to arguments
			if len(args) >= 2 {
				from, _ = strconv.Atoi(args[0])
				to, _ = strconv.Atoi(args[1])
			}

			// if 3 arguments passed set the increment
			if len(args) == 3 {
				increment, _ = strconv.Atoi(args[2])
			}

			// Increment must be greater than 0
			if increment == 0 {
				fmt.Println("seq: invalid increment '" + args[2] + "'")
				os.Exit(1)
			}

			// Make list of numbers
			numbers := []int{}

			// If increment is positive
			if increment > 0 {
				for i := from; i <= to; i += increment {
					numbers = append(numbers, i)
				}
			} else {
				for i := from; i >= to; i += increment {
					numbers = append(numbers, i)
				}
			}

			// print numbers separated by commas if the separator flag is set
			if separator, _ := cmd.Flags().GetBool("separator"); separator {
				// Convert numbers to strings
				numbersStr := []string{}
				for _, number := range numbers {
					numbersStr = append(numbersStr, strconv.Itoa(number))
				}

				fmt.Println(strings.Join(numbersStr, ","))
			} else {
				for _, number := range numbers {
					fmt.Println(number)
				}
			}
		},
	}

	// Add flags
	rootCmd.Flags().BoolP("separator", "s", false, "separate numbers by commas")

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

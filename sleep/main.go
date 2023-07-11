package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "sleep <duration>",
		Short: "Pause for the specified duration",
		Long:  "Pause for the specified duration. Optional unit suffix duration[suffix] (s for seconds, m for minutes, h for hours).",
		Run: func(cmd *cobra.Command, args []string) {
			// Check if all arguments were passed
			if len(args) == 0 {
				fmt.Println("sleep: missing operand")
				os.Exit(1)
			}

			// Check if the duration is valid
			durationArg := args[0]

			// if no suffix is specified, assume seconds
			if !strings.ContainsAny(string(durationArg[len(durationArg)-1]), "smh") {
				durationArg += "s"
			}

			duration, err := time.ParseDuration(durationArg)
			if err != nil {
				fmt.Println("sleep: invalid time interval '" + durationArg + "'")
				os.Exit(1)
			}

			// Sleep for the specified duration
			time.Sleep(duration)

		},
	}

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

/*
Copyright © 2024 Thar Htoo <iterdare.2977@gmail.com>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// timezoneCmd represents the timezone command
var timezoneCmd = &cobra.Command{
	Use:   "timezone",
	Short: "Get the current time in a given timezone",
	Long: `Get the current time in a given timezone.
				  This command takes one argument, the timezone you want to get the current time in.
				  It returns the current time in RFC1123 format.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timezone := args[0]
		formatFlag, _ := cmd.Flags().GetString("format")

		var date string
		location, _ := time.LoadLocation(timezone)
		if formatFlag != "" {
			date = time.Now().In(location).Format(formatFlag)
		} else {
			date = time.Now().In(location).Format(time.RFC3339)[:10]
		}
		fmt.Printf("Current date in %v: %v\n", timezone, date)
	},
}

func init() {
	rootCmd.AddCommand(timezoneCmd)
	// timezoneCmd.PersistentFlags().String("date", "", "returns the date in a time zone in a specified format")
	timezoneCmd.Flags().String("format", "", "Date for which to get the time (format: yyyy-mm-dd)")
}

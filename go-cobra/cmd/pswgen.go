/*
Copyright © 2024 Thar Htoo <iterdare.2977@gmail.com>
*/
package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

// pswgenCmd represents the pswgen command
var pswgenCmd = &cobra.Command{
	Use:   "pswgen",
	Short: "Generate random password",
	Long: `Generate random password with customizable options. 
For example:

go-cobra pswgen -l 12 -d -s`,
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")
		isDigits, _ := cmd.Flags().GetBool("digits")
		isSpecialChars, _ := cmd.Flags().GetBool("special-chars")

		charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

		if isDigits {
			charset += "0123456789"
		}

		if isSpecialChars {
			charset += "!@#$%^&*()_+-=[]{}|;:,.<>?~"
		}

		password := make([]byte, length)

		for i := range password {
			password[i] = charset[rand.Intn(len(charset))]
		}

		fmt.Println("Generating Password...")
		fmt.Println(string(password))
	},
}

func init() {
	rootCmd.AddCommand(pswgenCmd)
	pswgenCmd.Flags().IntP("length", "l", 8, "length of the generated password")
	pswgenCmd.Flags().BoolP("digits", "d", false, "Include digits in generated password")
	pswgenCmd.Flags().BoolP("special-chars", "s", false, "Include special characters in generated password")
}

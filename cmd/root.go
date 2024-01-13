/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goopenai",
	Short: "This is a CLI tool for interacting with OpenAI's API",
	Long:  ` `,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a CLI tool for interacting with OpenAI's API")
	},
}

var inputPrompt = &cobra.Command{
	Use:     "input",
	Short:   "input string prompt",
	Aliases: []string{"in"},
	// Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// prompt := prompts.InputPrompt()

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(inputPrompt)
}

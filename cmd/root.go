package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "dummygen",
	Short: "DummyGen generates dummy data from schema definitions",
	Long:  `DummyGen is a CLI tool that generates dummy data (SQL inserts, JSON, etc.) based on schema definitions.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to DummyGen! Use --help to see commands.")
	},
}

// Execute runs the root command
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var schemaFile string
var outputFile string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate dummy data from a schema",
	Long:  `Reads a schema file (JSON, XML in future) and generates dummy data as SQL insert statements.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Generating data from schema: %s -> %s\n", schemaFile, outputFile)
	},
}

func init() {
	// add flags

	generateCmd.Flags().StringVarP(&schemaFile, "schema", "s", "", "Path to schema file (required)")
	generateCmd.Flags().StringVarP(&outputFile, "output", "o", "output.sql", "Path to output file")

	// mark schema as required
	_ = generateCmd.MarkFlagRequired("schema")

	// register this command under root
	rootCmd.AddCommand(generateCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var memCmd = &cobra.Command{
	Use:   "mem",
	Short: "mem gets memory stats",
	Long:  "mem is a tool to get memory stats",
	Run: func(cmd *cobra.Command, args []string) {
		result := GetMemStats()
		fmt.Println(result)
		ExportToFile(result, exportFilePath)
	},
}

func init() {
	rootCmd.AddCommand(memCmd)
	memCmd.Flags().StringVarP(&exportFilePath, "export", "e", "", "Export to file")
}

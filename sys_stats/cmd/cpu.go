package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "cpu gets cpu stats",
	Long:  "cpu is a tool to get cpu stats",
	Run: func(cmd *cobra.Command, args []string) {
		result := GetCPUStats()
		fmt.Println(result)
		ExportToFile(result, exportFilePath)
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)
	cpuCmd.Flags().StringVarP(&exportFilePath, "export", "e", "", "Export to file")
}

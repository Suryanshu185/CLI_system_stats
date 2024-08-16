package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var diskCmd = &cobra.Command{
	Use:   "disk",
	Short: "disk gets disk stats",
	Long:  "disk is a tool to get disk stats",
	Run: func(cmd *cobra.Command, args []string) {
		result := GetDiskStats()
		fmt.Println(result)
		ExportToFile(result, exportFilePath)
	},
}

func init() {
	rootCmd.AddCommand(diskCmd)
	diskCmd.Flags().StringVarP(&exportFilePath, "export", "e", "", "Export to file")
}

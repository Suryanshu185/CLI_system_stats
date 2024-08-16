package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/spf13/cobra"
)

var exportFilePath string

func GetCPUUsage() string {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}
	return fmt.Sprintf("CPU Usage: %f%%", percent[0])
}

func GetMemUsage() string {
	v, err := mem.VirtualMemory()
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}
	return fmt.Sprintf("Memory Usage: %f%%", v.UsedPercent)
}

func GetDiskUsage() string {
	diskStats, err := disk.Usage("/")
	if err != nil {
		return fmt.Sprintf("Error: %s", err)
	}
	return fmt.Sprintf("Disk Usage: total: %d, used: %d, free: %d", diskStats.Total, diskStats.Used, diskStats.Free)
}

func GetCPUStats() string {
	return GetCPUUsage()
}

func GetDiskStats() string {
	return GetDiskUsage()
}

func GetMemStats() string {
	return GetMemUsage()
}

var statsCmd = &cobra.Command{
	Use:   "sys_stats",
	Short: "sys_stats gets system stats",
	Long:  "sys_stats is a tool to get system stats",
	Run: func(cmd *cobra.Command, args []string) {
		cpuUsage := GetCPUStats()
		memUsage := GetMemStats()
		diskUsage := GetDiskStats()

		result := strings.Join([]string{cpuUsage, memUsage, diskUsage}, "\n")
		fmt.Println(result)

		if exportFilePath != "" {
			ExportToFile(result, exportFilePath)
		}
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
	statsCmd.Flags().StringVarP(&exportFilePath, "export", "e", "", "Export to file")
}

func ExportToFile(result, filePath string) {
	if filePath != "" {
		formattedData := strings.ReplaceAll(result, " ", "\n")

		timeStamp := time.Now().Format("2006-01-02 15:04:05")
		logEntry := fmt.Sprintf("Time: %s\n%s", timeStamp, formattedData)

		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		defer file.Close()

		if _, err := file.WriteString(logEntry); err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Println("Data exported to file")
		}
	}
}

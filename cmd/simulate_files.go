package cmd

import (
	"edr-activity-generator/activity"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var simulateFilesCmd = &cobra.Command{
	Use:   "files",
	Short: "Simulate file creation/modification/deletion",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll("logs", os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}

		fmt.Printf("Simulating file activity. Output: %s\n", outputPath)
		if err := activity.PerformFileActivity(outputPath, outputFormat); err != nil {
			log.Fatalf("Error simulating file activity: %v", err)
		}
	},
}

func init() {
	simulateCmd.AddCommand(simulateFilesCmd)
}

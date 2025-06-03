package cmd

import (
	"edr-activity-generator/activity"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var simulateProcessCmd = &cobra.Command{
	Use:   "process",
	Short: "Simulate process creation",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll("logs", os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}

		fmt.Printf("Simulating process activity. Output: %s\n", outputPath)
		if err := activity.StartProcess(outputPath); err != nil {
			log.Fatalf("Error simulating process activity: %v", err)
		}
	},
}

func init() {
	simulateCmd.AddCommand(simulateProcessCmd)
}

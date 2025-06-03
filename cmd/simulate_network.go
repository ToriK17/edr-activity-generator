package cmd

import (
	"edr-activity-generator/activity"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var simulateNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "Simulate network traffic",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll("logs", os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}

		fmt.Printf("Simulating network activity. Output: %s\n", outputPath)
		if err := activity.SimulateNetworkActivity(outputPath, outputFormat); err != nil {
			log.Fatalf("Error simulating network HTTP/1.1 activity: %v", err)
		}
		if err := activity.SimulateHTTP2Activity(outputPath, outputFormat); err != nil {
			log.Fatalf("Error simulating HTTP/2 activity: %v", err)
		}
	},
}

func init() {
	simulateCmd.AddCommand(simulateNetworkCmd)
}

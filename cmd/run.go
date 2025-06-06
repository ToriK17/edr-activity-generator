package cmd

import (
	"edr-activity-generator/activity"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var outputPath string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Simulate endpoint activity and write a structured telemetry log",
	Run: func(cmd *cobra.Command, args []string) {
		outputDir := filepath.Dir(outputPath)
		err := os.MkdirAll(outputDir, 0755)
		// higher permissions needed here to enter custom paths and access files located inside.

		if err != nil {
			log.Fatalf("Failed to create log directory %q: %v", outputDir, err)
		}

		fmt.Println("Generating EDR test activity...")
		fmt.Printf("Logs will be written to %s\n", outputPath)

		if err := activity.StartProcess(outputPath, outputFormat); err != nil {
			log.Fatalf("Error generating process activity: %v", err)
		}

		if err := activity.PerformFileActivity(outputPath, outputFormat); err != nil {
			log.Fatalf("Error performing file activity: %v", err)
		}

		if err := activity.SimulateNetworkActivity(outputPath, outputFormat); err != nil {
			log.Fatalf("Error performing network activity: %v", err)
		}

		if err := activity.SimulateHTTP2Activity(outputPath, outputFormat); err != nil {
			log.Fatalf("Error performing HTTP/2 activity: %v", err)
		}

		fmt.Println("All activities completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&outputPath, "output", "o", "logs/activity_log.json", "Path to output log file")
}

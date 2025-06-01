package cmd

import (
	"edr-activity-generator/activity"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var outputPath string

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Simulate endpoint activity and write a structured telemetry log",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll("logs", os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}

		fmt.Println("Generating EDR test activity...")
		fmt.Printf("Logs will be written to %s\n", outputPath)

		if err := activity.StartProcess(outputPath); err != nil {
			log.Fatalf("Error generating process activity: %v", err)
		}

		if err := activity.PerformFileActivity(outputPath); err != nil {
			log.Fatalf("Error performing file activity: %v", err)
		}

		fmt.Println("All activities completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringVarP(&outputPath, "output", "o", "logs/activity_log.json", "Path to output log file")
}

package cmd

import (
	"edr-activity-generator/activity"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var simulateFilesCmd = &cobra.Command{
	Use:   "files",
	Short: "Simulate file creation / modification / deletion",
	Run: func(cmd *cobra.Command, args []string) {
		// Ensure logs directory exists
		if err := os.MkdirAll("logs", os.ModePerm); err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}

		start := time.Now()
		counted := 0

		for {
			if duration == 0 && counted >= count {
				break
			}
			if duration > 0 && time.Since(start) >= duration {
				break
			}

			fmt.Printf("Simulating file activity. Output: %s\n", outputPath)
			if err := activity.PerformFileActivity(outputPath, outputFormat); err != nil {
				log.Fatalf("Error simulating file activity: %v", err)
			}

			counted++
			if delay > 0 {
				time.Sleep(delay)
			}
		}
	},
}

func init() {
	simulateFilesCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of file-activity simulations to perform (ignored when --stream is set)")
	simulateFilesCmd.Flags().DurationVarP(&delay, "delay", "d", 0, "Delay between simulations (e.g. 500ms, 2s)")
	simulateFilesCmd.Flags().DurationVar(&duration, "stream", 0, "Continuously simulate for the specified duration (e.g. 30s)")
	simulateCmd.AddCommand(simulateFilesCmd)
}

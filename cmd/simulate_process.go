package cmd

import (
	"edr-activity-generator/activity"
	"fmt"
	"log"
	"os"
	"time"

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

		start := time.Now()
		counted := 0

		for {

			if duration == 0 && counted >= count {
				break
			}

			if duration > 0 && time.Since(start) >= duration {
				break
			}

			fmt.Printf("Simulating process activity. Output: %s\n", outputPath)
			if err := activity.StartProcess(outputPath, outputFormat); err != nil {
				log.Fatalf("Error simulating process activity: %v", err)
			}

			counted++
			if delay > 0 {
				time.Sleep(delay)
			}
		}
	},
}

func init() {
	simulateProcessCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of process simulations to perform")
	simulateProcessCmd.Flags().DurationVarP(&delay, "delay", "d", 0, "Delay between process simulations (e.g. 500ms, 2s)")
	simulateProcessCmd.Flags().DurationVar(&duration, "stream", 0, "Run simulations continuously for specified duration (e.g. 30s)")
	simulateCmd.AddCommand(simulateProcessCmd)
}

package cmd

import (
	"edr-activity-generator/activity"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var simulateNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "Simulate HTTP/1.1 and HTTP/2 traffic",
	Run: func(cmd *cobra.Command, args []string) {
		// Make sure the logs directory exists
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

			fmt.Printf("Simulating network activity. Output: %s\n", outputPath)

			if err := activity.SimulateNetworkActivity(outputPath, outputFormat); err != nil {
				log.Fatalf("Error simulating HTTP/1.1 activity: %v", err)
			}
			if err := activity.SimulateHTTP2Activity(outputPath, outputFormat); err != nil {
				log.Fatalf("Error simulating HTTP/2 activity: %v", err)
			}

			counted++
			if delay > 0 {
				time.Sleep(delay)
			}
		}
	},
}

func init() {
	simulateNetworkCmd.Flags().IntVarP(&count, "count", "c", 1, "Number of network simulations to perform (ignored when --stream is set)")
	simulateNetworkCmd.Flags().DurationVarP(&delay, "delay", "d", 0, "Delay between simulations (e.g. 500ms, 2s)")
	simulateNetworkCmd.Flags().DurationVar(&duration, "stream", 0, "Continuously simulate for the specified duration (e.g. 30s)")
	simulateCmd.AddCommand(simulateNetworkCmd)
}

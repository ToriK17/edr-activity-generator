package cmd

import (
	"edr-activity-generator/activity"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var simulateCmd = &cobra.Command{
	Use:   "simulate",
	Short: "Run individual types of activity simulation",
	Long:  `Simulate only one type of endpoint activity: process, files, or network.`,
}

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

var simulateFilesCmd = &cobra.Command{
	Use:   "files",
	Short: "Simulate file creation/modification/deletion",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll("logs", os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}

		fmt.Printf("Simulating file activity. Output: %s\n", outputPath)
		if err := activity.PerformFileActivity(outputPath); err != nil {
			log.Fatalf("Error simulating file activity: %v", err)
		}
	},
}

var simulateNetworkCmd = &cobra.Command{
	Use:   "network",
	Short: "Simulate network traffic",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.MkdirAll("logs", os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}

		fmt.Printf("Simulating network activity. Output: %s\n", outputPath)
		if err := activity.SimulateNetworkActivity(outputPath); err != nil {
			log.Fatalf("Error simulating network HTTP/1.1 activity: %v", err)
		}
		if err := activity.SimulateHTTP2Activity(outputPath); err != nil {
			log.Fatalf("Error simulating HTTP/2 activity: %v", err)
		}
	},
}

func init() {
	// Reuse the same output flag so all subcommands inherit --output
	simulateCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "logs/activity_log.json", "Path to output log file")

	simulateCmd.AddCommand(simulateProcessCmd)
	simulateCmd.AddCommand(simulateFilesCmd)
	simulateCmd.AddCommand(simulateNetworkCmd)

	rootCmd.AddCommand(simulateCmd)
}

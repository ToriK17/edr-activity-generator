package cmd

import (
	"github.com/spf13/cobra"
)

var simulateCmd = &cobra.Command{
	Use:   "simulate",
	Short: "Run individual types of activity simulation",
	Long:  `Simulate only one type of endpoint activity: process, files, or network.`,
}

func init() {
	// Share the --output flag across all simulate subcommands
	simulateCmd.PersistentFlags().StringVarP(&outputPath, "output", "o", "logs/activity_log.json", "Path to output log file")

	// Register simulate command with root
	rootCmd.AddCommand(simulateCmd)
}

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "edr-activity-generator",
	Short: "Simulate and log endpoint activity to validate EDR telemetry",
	Long: `edr-activity-generator is a CLI tool that triggers realistic endpoint activity 
		across supported platforms (Linux and macOS) and logs structured telemetry data 
	to help identify regressions in EDR agent output.

	This tool is designed for security validation use cases and logs:
		- Process creation
		- File creation, modification, and deletion

	Usage examples:
	edr-activity-generator run --output logs/activity_log.json
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// No global flags needed yet.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

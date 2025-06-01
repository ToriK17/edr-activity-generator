package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Remove the activity log file",
	Run: func(cmd *cobra.Command, args []string) {
		err := os.Remove("logs/activity_log.json")
		if err != nil && !os.IsNotExist(err) {
			log.Fatalf("Failed to remove log file: %v", err)
		}
		fmt.Println("🧹 Log file cleaned.")
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}

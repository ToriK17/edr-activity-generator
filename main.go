package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Define our first CLI command
	// Signature: func String(name string, value string, usage string) *string
	outputPath := flag.String("output", "logs/activity_log.json", "Path to output log file")
	flag.Parse()
	// make a dir logs and parent dir if needed and give with full read, write, and execute permissions
	// end program immediately if there are no logs, bail fast
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	fmt.Println("Generating EDR test activity...")
	fmt.Printf("Logs will be written to %s\n", *outputPath)

}

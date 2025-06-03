package main

// Import here so all the subcommands are mounted for simulate
import "edr-activity-generator/cmd"

func main() {
	cmd.Execute()
}

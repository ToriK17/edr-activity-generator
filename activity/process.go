package activity

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
	"time"
)

type ProcessLog struct {
	Timestamp   string `json:"timestamp"`
	Username    string `json:"username"`
	ProcessName string `json:"process_name"`
	CommandLine string `json:"command_line"`
	ProcessID   int    `json:"process_id"`
}

func StartProcess(outputPath string, format string, cmdArgs []string) error {
	// Tested with simple sleep 1 example, cross-platform command (linux and macOS)
	if len(cmdArgs) == 0 {
		return fmt.Errorf("no command provided")
	}
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	// Start the process
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start process: %w", err)
	}

	// Username that started the process
	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}
	// ensures it always has something like /bin/sleep
	resolvedPath, err := exec.LookPath(cmd.Args[0])
	if err != nil {
		return fmt.Errorf("failed to resolve process path: %w", err)
	}

	// Create log entry
	logEntry := ProcessLog{
		Timestamp:   time.Now().Format(time.RFC3339),
		Username:    currentUser.Username,
		ProcessName: resolvedPath,
		CommandLine: fmt.Sprintf("%s %s", cmd.Path, strings.Join(cmd.Args[1:], " ")),
		ProcessID:   cmd.Process.Pid,
	}

	// Wait for the command to finish, important if you want to avoid zombie processes (on Unix-like systems).
	err = cmd.Wait()
	if err != nil {
		return fmt.Errorf("process exited with error: %w", err)
	}

	// Open the log file for appending
	// 0644 = read/write for owner, read-only for group and others
	file, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	// Encode and write log
	err = logWriter(logEntry, outputPath, format)
	if err != nil {
		return fmt.Errorf("failed to write log: %w", err)
	}

	fmt.Println("Process activity successfully logged.")
	return nil
}

package activity

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

type FileLog struct {
	Timestamp   string `json:"timestamp"`
	FilePath    string `json:"file_path"`
	Action      string `json:"action"`
	Username    string `json:"username"`
	ProcessName string `json:"process_name"`
	CommandLine string `json:"command_line"`
	ProcessID   int    `json:"process_id"`
}

func PerformFileActivity(outputPath, format string) error {
	filePath := "test_output.txt"
	fullPath, _ := filepath.Abs(filePath)

	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	processID := os.Getpid()
	processName := os.Args[0]
	commandLine := strings.Join(os.Args, " ")

	logAction := func(action string) error {
		logEntry := FileLog{
			Timestamp:   time.Now().Format(time.RFC3339),
			FilePath:    fullPath,
			Action:      action,
			Username:    currentUser.Username,
			ProcessName: processName,
			CommandLine: commandLine,
			ProcessID:   processID,
		}

		return writeLog(logEntry, outputPath, format)
	}

	// Create a file
	err = os.WriteFile(fullPath, []byte("Initial content\n"), 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}

	if err := logAction("create"); err != nil {
		return fmt.Errorf("failed to log create activity: %w", err)
	}

	// Modify the file
	f, err := os.OpenFile(fullPath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file for modification: %w", err)
	}
	_, err = f.WriteString("Appended line\n")
	if err != nil {
		return fmt.Errorf("failed to modify file: %w", err)
	}
	f.Close()

	if err := logAction("modify"); err != nil {
		return fmt.Errorf("failed to log modifying activity: %w", err)
	}

	// Delete the file
	err = os.Remove(fullPath)
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}

	if err := logAction("delete"); err != nil {
		return fmt.Errorf("failed to log deleting activity: %w", err)
	}

	fmt.Println("File activity logged.")
	return nil
}

package activity

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/user"
	"strings"
	"time"
)

type NetworkLog struct {
	Timestamp   string `json:"timestamp"`
	Username    string `json:"username"`
	SourceAddr  string `json:"source_address"`
	Destination string `json:"destination_address"`
	Protocol    string `json:"protocol"`
	BytesSent   int    `json:"bytes_sent"`
	ProcessName string `json:"process_name"`
	CommandLine string `json:"command_line"`
	ProcessID   int    `json:"process_id"`
}

func SimulateNetworkActivity(outputPath string) error {
	const defaultTarget = "example.com:80"

	conn, err := net.Dial("tcp", defaultTarget)
	if err != nil {
		return fmt.Errorf("failed to open TCP connection: %w", err)
	}
	defer conn.Close()

	message := "HEAD / HTTP/1.1\r\nHost: example.com\r\n\r\n"
	// \r\n\r\n = end of all headers, start of body needed specifically for HTTP/1.1
	bytesSent, err := conn.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("failed to send data: %w", err)
	}

	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	logEntry := NetworkLog{
		Timestamp:   time.Now().Format(time.RFC3339),
		Username:    currentUser.Username,
		SourceAddr:  conn.LocalAddr().String(),
		Destination: conn.RemoteAddr().String(),
		Protocol:    "tcp",
		BytesSent:   bytesSent,
		ProcessName: os.Args[0],
		CommandLine: strings.Join(os.Args, " "),
		ProcessID:   os.Getpid(),
	}

	// Write to log file
	file, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(logEntry); err != nil {
		return fmt.Errorf("failed to write network log: %w", err)
	}

	fmt.Println("Network activity logged.")
	return nil
}

package activity

import (
	"fmt"
	"net"
	"net/http"
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

func SimulateNetworkActivity(outputPath string, format string) error {
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
		Timestamp:   time.Now().Format(time.RFC3339), // ensures the timestamp is in ISO 8601, human readable/ machine parsable
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

	err = writeLog(logEntry, outputPath, format)
	if err != nil {
		return fmt.Errorf("failed to write network log: %w", err)
	}

	fmt.Println("HTTP/1.1 Network activity logged.")
	return nil
}

func SimulateHTTP2Activity(outputPath string, format string) error {
	const targetURL = "https://nghttp2.org"

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, err := http.NewRequest("HEAD", targetURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP/2 request: %w", err)
	}

	start := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to perform HTTP/2 request: %w", err)
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)

	// Check negotiated protocol
	connState := resp.TLS
	protocol := "unknown"
	if connState != nil && len(connState.NegotiatedProtocol) > 0 {
		protocol = connState.NegotiatedProtocol
	}

	currentUser, err := user.Current()
	if err != nil {
		return fmt.Errorf("failed to get current user: %w", err)
	}

	logEntry := NetworkLog{
		Timestamp:   time.Now().Format(time.RFC3339),
		Username:    currentUser.Username,
		SourceAddr:  "-", // http.Client doesn't expose local address details like source IP/port it's abstracted by the transport layer
		Destination: targetURL,
		Protocol:    protocol,
		BytesSent:   0, // tracking sent byte count requires lower-level instrumentation (e.g. httptrace hooks or custom Transport)
		ProcessName: os.Args[0],
		CommandLine: strings.Join(os.Args, " "),
		ProcessID:   os.Getpid(),
	}

	file, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer file.Close()

	err = writeLog(logEntry, outputPath, format)
	if err != nil {
		return fmt.Errorf("failed to write HTTP/2 network log: %w", err)
	}

	fmt.Printf("HTTP/2 activity logged (%s) in %v\n", protocol, elapsed)
	return nil
}

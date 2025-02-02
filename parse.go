package main

import (
	"regexp"
	"strings"
	"time"
	"fmt"
)

type LogEntry struct {
	Timestamp time.Time `json:"timestamp"`
	Hostname  string    `json:"hostname"`
	Source    string    `json:"source"`
	Message   string    `json:"message"`
	Severity  string    `json:"severity"`
}

func processLogData(data string, includePatterns []*regexp.Regexp, file string) []LogEntry {
	lines := strings.Split(data, "\n")
	return aggregateLogs(lines, includePatterns, file)
}

func aggregateLogs(logLines []string, includePatterns []*regexp.Regexp, file string) []LogEntry {
	ip, err := getServerIP()
	if err != nil {
		fmt.Printf("Error getting server IP: %s\n", err)
		ip = "unknown" // Fallback if IP retrieval fails
	}

	var entries []LogEntry
	for _, line := range logLines {
		if len(line) == 0 {
			continue
		}

		for _, pattern := range includePatterns {
			severity := "INFO"
			if pattern.MatchString(line) {
				severity = "WARNING"
				break
			}

			entry := LogEntry{
				Timestamp: time.Now(),
				Hostname:  ip, // Adjust as needed
				Source:    file, // Adjust
				Message:   line,
				Severity:  severity, // Default severity
			}
			entries = append(entries, entry)
		}
	}
	return entries
}

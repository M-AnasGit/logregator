package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Application running")

	// Define flags
	key := flag.String("k", "", "A key that must be provided")
	files := flag.String("f", "", "Comma-separated list of files to include")

	// Parse flags
	flag.Parse()

	// Ensure key is provided
	if *key == "" {
		fmt.Fprintln(os.Stderr, "Error: -k argument is required")
		os.Exit(1)
	}

	// Process log files from -f argument
	logFiles := []string{}
	if *files != "" {
		logFiles = append(logFiles, splitCommaSeparated(*files)...)
	}

	connectToSocket(key)

	// Add default log files if none provided via -f
	defaultLogFiles := []string{"/var/log/syslog", "/var/log/auth.log"}
	if len(logFiles) == 0 {
		logFiles = append(logFiles, defaultLogFiles...)
	}

	// Load include patterns
	includeFolder := "./include.d"
	includePatterns, err := loadRegexPatterns(includeFolder)
	fatal_check(err)

	// Watch the log files
	watchLogFiles(logFiles, includePatterns)
}

// Helper function to split comma-separated string into slice
func splitCommaSeparated(input string) []string {
	var result []string
	for _, item := range strings.Split(input, ",") {
		result = append(result, strings.TrimSpace(item))
	}
	return result
}

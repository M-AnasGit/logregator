package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"os"
	"regexp"
)

func handleLogEntries(entries []LogEntry) {
	// Check if there are more than 5 entries
	if len(entries) > 5 {
		// If so, take only the latest 5 entries
		entries = entries[len(entries)-5:]
	}

	// Print the selected log entries
	for _, entry := range entries {
		fmt.Printf("Log Entry: %+v\n", entry) // Use %+v for more detailed output
	}
}

func watchLogFiles(logFiles []string, includePatterns []*regexp.Regexp) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating file watcher:", err)
		return
	}
	defer watcher.Close()

	for _, file := range logFiles {
		err = watcher.Add(file)
		if err != nil {
			fmt.Printf("Error watching file %s: %s\n", file, err)
		}
	}

	fmt.Println("Monitoring logs:", logFiles)

	// Process logs at program startup
	for _, file := range logFiles {
		data, err := os.ReadFile(file)
		if err != nil {
			fmt.Printf("Error reading log file %s: %s\n", file, err)
			continue
		}
		entries := processLogData(string(data), includePatterns, file)
		handleLogEntries(entries) // Call the new handler
	}

	// Watch for changes in log files
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					data, err := os.ReadFile(event.Name)
					if err != nil {
						fmt.Printf("Error reading log file %s: %s\n", event.Name, err)
						continue
					}

					// Process log data and aggregate entries
					entries := processLogData(string(data), includePatterns, event.Name)
					handleLogEntries(entries) // Call the new handler
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				fmt.Println("Error:", err)
			}
		}
	}()

	// Block forever
	select {}
}

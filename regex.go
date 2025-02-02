package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func loadRegexPatterns(folderPath string) ([]*regexp.Regexp, error) {
	var patterns []*regexp.Regexp
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %s: %w", path, err)
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("error opening file %s: %w", path, err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if len(line) > 0 && line[0] != '#' {
					r, err := regexp.Compile(line)
					if err != nil {
						fmt.Printf("Invalid regex in file %s: %s\nError: %s\n", path, line, err)
						continue // Skip invalid regex and move to the next line
					}
					patterns = append(patterns, r)
				}
			}

			if err := scanner.Err(); err != nil {
				return fmt.Errorf("error reading file %s: %w", path, err)
			}
		}
		return nil
	})

	return patterns, err
}
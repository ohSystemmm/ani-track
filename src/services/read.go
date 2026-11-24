package services

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func LoadEntries() ([]Entry, error) {
	dataPath := GetDataFilePath()

	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		if err := EnsureSeedData(); err != nil {
			return nil, err
		}
	}

	content, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, fmt.Errorf("error reading %s: %w", dataPath, err)
	}

	var entries []Entry

	trimmed := strings.TrimSpace(string(content))
	if strings.HasSuffix(dataPath, ".data") || (len(trimmed) > 0 && trimmed[0] != '[') {
		entries, err = parseLegacyData(string(content))
		if err != nil {
			return nil, fmt.Errorf("error parsing legacy data: %w", err)
		}
	} else {
		if len(trimmed) > 0 {
			if err := json.Unmarshal([]byte(trimmed), &entries); err != nil {
				entries, err = parseLegacyData(string(content))
				if err != nil {
					return nil, fmt.Errorf("error unmarshaling json entries: %w", err)
				}
			}
		}
	}

	for i := range entries {
		entries[i].Index = i + 1
	}

	return entries, nil
}

func GetEntry(index int) (*Entry, error) {
	entries, err := LoadEntries()
	if err != nil {
		return nil, err
	}

	for _, e := range entries {
		if e.Index == index {
			return &e, nil
		}
	}

	return nil, fmt.Errorf("entry with index %d not found", index)
}

func GetHistory() ([]Entry, error) {
	entries, err := LoadEntries()
	if err != nil {
		return nil, err
	}

	var history []Entry
	for _, e := range entries {
		if e.IsFinished() {
			history = append(history, e)
		}
	}
	return history, nil
}

func GetWatchlist() ([]Entry, error) {
	entries, err := LoadEntries()
	if err != nil {
		return nil, err
	}

	var watchlist []Entry
	for _, e := range entries {
		if !e.IsFinished() {
			watchlist = append(watchlist, e)
		}
	}
	return watchlist, nil
}

func parseLegacyData(content string) ([]Entry, error) {
	scanner := bufio.NewScanner(strings.NewReader(content))
	var entries []Entry
	current := Entry{}

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			if current.Title != "" {
				entries = append(entries, current)
				current = Entry{}
			}
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.ToLower(strings.TrimSpace(parts[0]))
		val := strings.TrimSpace(parts[1])

		switch key {
		case "title":
			current.Title = val
		case "category":
			current.Category = val
		case "rating":
			current.Rating = val
		case "started":
			current.Started = val
		case "finished":
			current.Finished = val
		case "status":
			current.Status = val
		case "progress":
			current.Progress = val
		case "notes":
			current.Notes = val
		}
	}

	if current.Title != "" {
		entries = append(entries, current)
	}

	return entries, scanner.Err()
}

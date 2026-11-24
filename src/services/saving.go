package services

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func SaveAllEntries(entries []Entry) error {
	for i := range entries {
		entries[i].Index = i + 1
	}

	data, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling entries: %w", err)
	}

	dataPath := GetDataFilePath()
	if filepath.Ext(dataPath) == ".data" {
		dataPath = LocalConfigPath
	}

	if err := os.MkdirAll(filepath.Dir(dataPath), 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	if err := os.WriteFile(dataPath, data, 0644); err != nil {
		return fmt.Errorf("error writing entries to %s: %w", dataPath, err)
	}

	return nil
}

func AddEntry(entry Entry) error {
	entries, err := LoadEntries()
	if err != nil {
		entries = []Entry{}
	}

	entry.Index = len(entries) + 1
	entries = append(entries, entry)
	return SaveAllEntries(entries)
}

func UpdateEntry(index int, updated Entry) error {
	entries, err := LoadEntries()
	if err != nil {
		return err
	}

	found := false
	for i := range entries {
		if entries[i].Index == index {
			updated.Index = index
			entries[i] = updated
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("entry with index %d not found", index)
	}

	return SaveAllEntries(entries)
}

func DeleteEntry(index int) error {
	entries, err := LoadEntries()
	if err != nil {
		return err
	}

	newEntries := make([]Entry, 0, len(entries))
	for _, e := range entries {
		if e.Index != index {
			newEntries = append(newEntries, e)
		}
	}

	return SaveAllEntries(newEntries)
}

func Save(title, category, rating, started, finished, status, notes string) error {
	return AddEntry(Entry{
		Title:    title,
		Category: category,
		Rating:   FormatRating(rating),
		Started:  started,
		Finished: finished,
		Status:   status,
		Notes:    notes,
	})
}

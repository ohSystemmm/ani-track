package services

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var (
	LocalConfigPath = filepath.Join("config", "Ani-Track.json")
	LegacyDataPath  = filepath.Join("config", "Ani-Track.data")
)

func DefaultInitialEntries() []Entry {
	return []Entry{
		{
			Index:    1,
			Title:    "Frieren: Beyond Journey's End",
			Category: "Anime",
			Rating:   "98/100",
			Started:  "Sep 2023",
			Finished: "Mar 2024",
			Status:   "Finished",
			Progress: "28/28",
			Notes:    "Masterpiece of pacing, reflection, and fantasy worldbuilding.",
		},
		{
			Index:    2,
			Title:    "Steins;Gate",
			Category: "Anime",
			Rating:   "96/100",
			Started:  "Apr 2021",
			Finished: "May 2021",
			Status:   "Finished",
			Progress: "24/24",
			Notes:    "Peak sci-fi thriller. El Psy Kongroo.",
		},
		{
			Index:    3,
			Title:    "Berserk",
			Category: "Manga",
			Rating:   "99/100",
			Started:  "2019",
			Finished: "2023",
			Status:   "Finished",
			Progress: "Vol 41",
			Notes:    "Kentaro Miura's unforgettable dark fantasy epic.",
		},
		{
			Index:    4,
			Title:    "One Piece",
			Category: "Anime",
			Rating:   "93/100",
			Started:  "2022",
			Finished: "Ongoing",
			Status:   "In Watch",
			Progress: "Ep 1110",
			Notes:    "The adventure never ends! Egghead arc is peak.",
		},
		{
			Index:    5,
			Title:    "Chainsaw Man",
			Category: "Manga",
			Rating:   "90/100",
			Started:  "2023",
			Finished: "Ongoing",
			Status:   "In Read",
			Progress: "Ch 165",
			Notes:    "Part 2 continues to be wild and unpredictable.",
		},
		{
			Index:    6,
			Title:    "Dungeon Meshi",
			Category: "Anime",
			Rating:   "92/100",
			Started:  "Jan 2024",
			Finished: "Ongoing",
			Status:   "In Watch",
			Progress: "Ep 24/24",
			Notes:    "Delicious cooking adventures combined with great lore.",
		},
		{
			Index:    7,
			Title:    "Solo Leveling",
			Category: "Novel",
			Rating:   "88/100",
			Started:  "2024",
			Finished: "Ongoing",
			Status:   "Plan to Watch",
			Progress: "Ch 50",
			Notes:    "Exciting power fantasy and dungeon raids.",
		},
	}
}

func GetDataFilePath() string {
	if _, err := os.Stat(LocalConfigPath); err == nil {
		return LocalConfigPath
	}

	if _, err := os.Stat(LegacyDataPath); err == nil {
		return LegacyDataPath
	}

	homeDir, err := os.UserHomeDir()
	if err == nil {
		xdgPath := filepath.Join(homeDir, ".local", "state", "ani-track", "Ani-Track.json")
		if _, err := os.Stat(xdgPath); err == nil {
			return xdgPath
		}
	}

	return LocalConfigPath
}

func ResolveDirectory(path string) (string, error) {
	if strings.HasPrefix(path, "~/") || strings.HasPrefix(path, "~\\") || path == "~" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		if path == "~" {
			path = homeDir
		} else {
			path = filepath.Join(homeDir, path[2:])
		}
	}
	return filepath.Clean(path), nil
}

func CopyFile(src, dst string) error {
	inFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("error opening source file: %w", err)
	}
	defer inFile.Close()

	if err := os.MkdirAll(filepath.Dir(dst), 0755); err != nil {
		return err
	}

	outFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("error creating destination file: %w", err)
	}
	defer outFile.Close()

	if _, err := io.Copy(outFile, inFile); err != nil {
		return fmt.Errorf("error copying file: %w", err)
	}
	return nil
}

func EnsureSeedData() error {
	dataPath := GetDataFilePath()
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		initial := DefaultInitialEntries()
		data, err := json.MarshalIndent(initial, "", "  ")
		if err != nil {
			return err
		}
		if err := os.MkdirAll(filepath.Dir(dataPath), 0755); err != nil {
			return err
		}
		return os.WriteFile(dataPath, data, 0644)
	}
	return nil
}

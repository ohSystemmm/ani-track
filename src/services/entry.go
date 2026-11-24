package services

import (
	"fmt"
	"strings"
)

type Entry struct {
	Index    int    `json:"Index"`
	Title    string `json:"Title"`
	Category string `json:"Category"`
	Rating   string `json:"Rating"`
	Started  string `json:"Started"`
	Finished string `json:"Finished"`
	Status   string `json:"Status"`
	Progress string `json:"Progress"`
	Notes    string `json:"Notes"`
}

func (e Entry) ToString() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Index: %d, Title: %s, Category: %s, Rating: %s, Started: %s, Finished: %s, Status: %s, Progress: %s, Notes: %s",
		e.Index, e.Title, e.Category, e.Rating, e.Started, e.Finished, e.Status, e.Progress, e.Notes))
	return sb.String()
}

func (e Entry) IsFinished() bool {
	return strings.EqualFold(strings.TrimSpace(e.Status), "Finished")
}

func FormatRating(val string) string {
	val = strings.TrimSpace(val)
	if val == "" {
		return "N/A"
	}
	if strings.Contains(val, "/") {
		return val
	}
	return val + "/100"
}

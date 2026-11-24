package services

func Reload() ([]Entry, error) {
	return LoadEntries()
}

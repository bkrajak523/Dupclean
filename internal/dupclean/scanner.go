package dupclean

import (
	"os"
	"path/filepath"
)

func ScanDuplicates(root string) map[string][]string {
	filesByHash := make(map[string][]string)

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		hash, err := HashFile(path)
		if err == nil {
			filesByHash[hash] = append(filesByHash[hash], path)
		}
		return nil
	})

	// Keep only duplicates
	duplicates := make(map[string][]string)
	for hash, paths := range filesByHash {
		if len(paths) > 1 {
			duplicates[hash] = paths
		}
	}
	return duplicates
}

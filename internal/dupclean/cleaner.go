package dupclean

import (
	"fmt"
	"os"
	"path/filepath"
)

// Scan scans the given directory and returns a map of hash -> files
func Scan(root string) map[string][]string {
	dupes := make(map[string][]string)

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		hash, err := HashFile(path) // uses hasher.go
		if err != nil {
			fmt.Println("error hashing:", path, err)
			return nil
		}

		dupes[hash] = append(dupes[hash], path)
		return nil
	})

	return dupes
}

// DeleteFiles keeps first file, deletes rest
func DeleteFiles(files []string, dryRun bool) int {
	deleted := 0
	if len(files) <= 1 {
		return deleted
	}

	original := files[0]
	fmt.Printf("✅ Keeping: %s\n", original)

	for _, f := range files[1:] {
		if dryRun {
			fmt.Printf("🗑️ Would delete: %s\n", f)
		} else {
			err := os.Remove(f)
			if err != nil {
				fmt.Println("❌ Error deleting:", f, err)
			} else {
				fmt.Printf("🗑️ Deleted: %s\n", f)
				deleted++
			}
		}
	}
	return deleted
}

// Clean scans and deletes duplicates automatically
func Clean(root string, dryRun bool) {
	dupes := Scan(root)
	totalDeleted := 0

	for _, files := range dupes {
		if len(files) > 1 {
			totalDeleted += DeleteFiles(files, dryRun)
		}
	}

	fmt.Printf("\n✅ Cleanup complete. Deleted %d duplicate files.\n", totalDeleted)
}

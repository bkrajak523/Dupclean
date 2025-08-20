package dupclean

import (
	"encoding/json"
	"fmt"
	"os"
)

func PrintResults(results map[string][]string, asJSON bool) {
	if asJSON {
		if len(results) == 0 {
			// print empty JSON array instead of message
			fmt.Fprint(os.Stdout, "[]")
			return
		}
		enc := json.NewEncoder(os.Stdout)
		enc.SetIndent("", "  ")
		if err := enc.Encode(results); err != nil {
			fmt.Fprintf(os.Stderr, "error encoding JSON: %v\n", err)
		}
	} else {
		if len(results) == 0 {
			fmt.Println("✅ No duplicate files found.")
			return
		}
		fmt.Println("🔍 Duplicate files found:")
		for hash, files := range results {
			if len(files) > 1 {
				fmt.Printf("Hash: %s\n", hash)
				for _, f := range files {
					fmt.Printf("  %s\n", f)
				}
			}
		}
	}
}

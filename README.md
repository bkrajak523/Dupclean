# DupClean
A fast and safe CLI tool to detect and remove duplicate files from your system.

Overview / Description:-
DupClean scans directories for duplicate files based on **file content**, not filenames. 
It efficiently detects duplicates using a **two-step approach**:  

1. Group files by **size** (quickly eliminates non-duplicates).  
2. Compute **SHA256 hash** for files with same size to ensure content matches.  

The tool supports:-
- Dry-run mode for previewing duplicates.
- Automatic deletion of duplicates, keeping one original file.
- Recursive scanning of directories.


Features:-

- Scan directories for duplicate files.
- Detect duplicates based on content using SHA256 hashing.
- Dry-run mode: preview duplicates before deleting.
- Auto-clean mode: delete duplicates while keeping one copy.
- Cross-platform: works on Linux/macOS.
- Modular and extensible code (Go + Docker ready).


Installation:-
# Clone repository
git clone https://github.com/yourusername/dupclean.git
cd dupclean


How to Use:-
# Build the CLI
go build -o dupclean ./cmd/dupclean
# Scan a directory for duplicates
./dupclean scan ~/Downloads

# Dry-run clean (preview duplicates)
./dupclean clean ~/Downloads

# Auto-clean (delete duplicates automatically)
./dupclean clean ~/Downloads --auto


How it works:-
Scan directory -> Group by file size -> Hash same-size files -> Group by hash -> Keep one, delete rest

Techstacks:-
- Written in Go (Golang) for speed and concurrency.
- Uses Go’s `crypto/sha256` for content hashing.
- CLI powered by Cobra library.
- Modular design: `hasher.go`, `cleaner.go`, `prompt.go`.

Future Improvements:-
- Parallel hashing for very large directories.
- Cache hashes to avoid recomputation.
- GUI or web interface for easier visualization.
- Configurable exclusion patterns (ignore folders/files).



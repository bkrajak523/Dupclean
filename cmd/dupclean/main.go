package main

import (
	"fmt"
	"os"
	"example.com/dupclean/internal/dupclean"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "dupclean",
		Short: "Scan and clean duplicate files",
	}

	var auto bool

	var scanCmd = &cobra.Command{
		Use:   "scan [path]",
		Short: "Scan a directory for duplicates",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			dupes := dupclean.Scan(path)
			for _, files := range dupes {
				if len(files) > 1 {
					fmt.Println("Duplicate group:")
					for _, f := range files {
						fmt.Println("   ", f)
					}
				}
			}
		},
	}

	var cleanCmd = &cobra.Command{
		Use:   "clean [path]",
		Short: "Clean duplicates in a directory",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			dupclean.Clean(path, !auto)
		},
	}

	cleanCmd.Flags().BoolVar(&auto, "auto", false, "Automatically delete duplicates")

	rootCmd.AddCommand(scanCmd)
	rootCmd.AddCommand(cleanCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

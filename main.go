package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// const tempFilePath = "/tmp/relapath_dir"
const tempFilePath = "./test_folder/relapath_dir"

func main() {
	args := os.Args[1:]

	switch {
	case len(args) == 1 && args[0] == "set":
		setCurrentDirectory()
	case len(args) == 1 && args[0] == "calc":
		calculateAndPrintRelativePath()
	default:
		fmt.Println("Usage: relapath [set|calc]")
	}
}

func setCurrentDirectory() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	err = os.WriteFile(tempFilePath, []byte(cwd), 0644)
	if err != nil {
		fmt.Println("Error writing to temp file:", err)
	}
}

func calculateAndPrintRelativePath() {
	firstPathBytes, err := os.ReadFile(tempFilePath)
	if err != nil {
		fmt.Println("Error reading from temp file:", err)
		return
	}

	firstPath := string(firstPathBytes)
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}

	relativePath, err := filepath.Rel(firstPath, cwd)
	if err != nil {
		fmt.Println("Error calculating relative path:", err)
		return
	}

	fmt.Println("Relative path from first to second:", relativePath)
}

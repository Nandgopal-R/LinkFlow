package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func EnsureBlogFileExists() (string, error) {
	dirPath := os.Getenv("FILE_PATH")
	if dirPath == "" {
		return "", fmt.Errorf("FILE_PATH environment variable not set")
	}

	// Check if the directory exists.
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		fmt.Printf("Directory does not exist, creating: %s\n", dirPath)
		// Create the directory
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			return "", fmt.Errorf("failed to create directory: %w", err)
		}
	}

	// Construct the full path to the blogs.txt file.
	filePath := filepath.Join(dirPath, "blogs.txt")

	// Check if the file exists.
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		fmt.Printf("File does not exist, creating: %s\n", filePath)
		// Create an empty file.
		file, err := os.Create(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to create file: %w", err)
		}
		file.Close()
	}

	fmt.Println("File exists at:", filePath)

	return filePath, nil
}

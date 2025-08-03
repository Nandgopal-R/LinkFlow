package watcher

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pmezard/go-difflib/difflib"
	"log"
	"os"
)

func DiffFinder(oldContent string, newContent string) string {
	diff := difflib.UnifiedDiff{
		A:        difflib.SplitLines(oldContent),
		B:        difflib.SplitLines(newContent),
		FromFile: "Original",
		ToFile:   "Current",
		Context:  0,
	}
	text, err := difflib.GetUnifiedDiffString(diff)
	if err != nil {
		log.Fatal(err)
	}
	return text
}

func FileWatcher(watcher *fsnotify.Watcher, file_addr string) {
	deletedLines := make([]string, 0)
	addedLines := make([]string, 0)
	initialSnapshot, err := os.ReadFile(file_addr)
	if err != nil {
		log.Fatal(err)
	}
	for { // Creates an infinte loop to continuously watch for events
		select { // Select statement to handle multiple channel operations
		case event, ok := <-watcher.Events:
			if !ok {
				log.Println("Error reading from watcher events channel")
				return
			}
			fmt.Println("Event:", event) // Print the event details
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Println("File modified:", event.Name)  // Notify if a file is modified
				newSnapshot, err := os.ReadFile(file_addr) // Read the new newContent
				if err != nil {
					log.Println("Error reading file:", err) // Log error if reading fails
					continue
				}
				diff := DiffFinder(string(initialSnapshot), string(newSnapshot)) // Get the diff between old and new newContent
				diffSlices := difflib.SplitLines(diff)                           // Split the diff into lines
				for _, line := range diffSlices {
					if len(line) > 0 && line[0] == '-' && line[1] != '-' {
						deletedLines = append(deletedLines, line[1:]) // Collect deleted lines
					}
					if len(line) > 0 && line[0] == '+' && line[1] != '+' {
						addedLines = append(addedLines, line[1:]) // Collect added lines
					}
				}
				// deletedLines = diff
				fmt.Println("Diff:\n", diff)                // Print the diff
				fmt.Println("Deleted Lines:", deletedLines) // Print deleted lines
				fmt.Println("Added Lines:", addedLines)     // Print added lines
				initialSnapshot = newSnapshot               // Update the initial snapshot to the new content
				deletedLines = make([]string, 0)            // Reset deleted lines
				addedLines = make([]string, 0)              // Reset added lines
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				log.Println("Error reading from watcher errors channel")
				return
			}
			log.Println("Error:", err)
		}
	}
}

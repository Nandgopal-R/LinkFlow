package watcher

import (
	"context"
	"fmt"
	db "github.com/Nandgopal-R/LinkFLow/db/gen"
	"github.com/Nandgopal-R/LinkFLow/internal/util"
	"github.com/fsnotify/fsnotify"
	"github.com/jackc/pgx/v5"
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

func FileWatcher(watcher *fsnotify.Watcher, file_addr string, conn *pgx.Conn) {
	q := db.New(conn)

	deletedLines := make([]string, 0)
	addedLines := make([]string, 0)
	initialSnapshot, err := os.ReadFile(file_addr)
	if err != nil {
		log.Fatal(err)
	}
	for { // Infinte loop
		select { // Select statement to handle multiple channel operations
		case event, ok := <-watcher.Events:
			if !ok {
				log.Println("Error reading from watcher events channel")
				return
			}
			fmt.Println("Event:", event)
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Println("File modified:", event.Name)
				newSnapshot, err := os.ReadFile(file_addr)
				if err != nil {
					log.Println("Error reading file:", err) // Log error if reading fails
					continue
				}

				diff := DiffFinder(string(initialSnapshot), string(newSnapshot))
				diffSlices := difflib.SplitLines(diff)
				for _, line := range diffSlices {
					title, blog_url, desc := util.SplitString(line[1:])
					if len(line) > 0 && line[0] == '-' && line[1] != '-' {
						q.DeleteBlogQuery(context.Background(), blog_url)
						// deletedLines = append(deletedLines, line[1:]) // Collect deleted lines
					}
					if len(line) > 0 && line[0] == '+' && line[1] != '+' {
						blog := db.InsertBlogQueryParams{
							Title:       title,
							BlogUrl:     blog_url,
							Description: desc,
						}
						q.InsertBlogQuery(context.Background(), blog)
						addedLines = append(addedLines, line[1:]) // Collect added lines
					}
				}

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

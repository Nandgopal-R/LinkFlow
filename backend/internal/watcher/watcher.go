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
	"time"
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
			if event.Op&fsnotify.Write == fsnotify.Write || event.Op&fsnotify.Rename == fsnotify.Rename {
				fmt.Println("File modified:", event.Name)
				newSnapshot, err := os.ReadFile(file_addr)
				if err != nil {
					// If the file doesn't exist, it's likely our race condition.
					if os.IsNotExist(err) {
						// Wait 100 milliseconds for the editor to finish the rename.
						time.Sleep(100 * time.Millisecond)
						// Try reading the file again.
						newSnapshot, err = os.ReadFile(file_addr)
					}

					// If there's still an error after the retry, then skip this event.
					if err != nil {
						log.Printf("Error reading file after retry: %v", err)
						continue
					}
				}
				diff := DiffFinder(string(initialSnapshot), string(newSnapshot))
				diffSlices := difflib.SplitLines(diff)
				for _, line := range diffSlices {
					if len(line) > 0 && line[0] == '-' && line[1] != '-' {
						_, blog_url, _ := util.SplitString(line[1:])
						q.DeleteBlogQuery(context.Background(), blog_url)
					}
					if len(line) > 0 && line[0] == '+' && line[1] != '+' {
						title, blog_url, desc := util.SplitString(line[1:])
						blog := db.InsertBlogQueryParams{
							Title:       title,
							BlogUrl:     blog_url,
							Description: desc,
						}
						q.InsertBlogQuery(context.Background(), blog)
					}
				}
				initialSnapshot = newSnapshot // Update the initial snapshot to the new content
				err = watcher.Add(file_addr)  // Re-add the file to the watcher
				if err != nil {
					log.Printf("Error re-adding file to watcher: %v", err)
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				log.Println("Watcher errors channel closed.")
				return
			}
			log.Println("Error:", err)
		}
	}
}

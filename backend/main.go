package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Nandgopal-R/LinkFLow/internal/util"
	"github.com/Nandgopal-R/LinkFLow/internal/watcher"
	"github.com/fsnotify/fsnotify"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	filepath, err := util.EnsureBlogFileExists()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ensuring blog file exists: %v\n", err)
		os.Exit(1)
	}

	NewWatcher, err := fsnotify.NewWatcher() // Create a new file watcher
	if err != nil {
		log.Fatal(err)
	}
	defer NewWatcher.Close()

	err = NewWatcher.Add(filepath)
	if err != nil {
		log.Fatal(err)
	}

	go watcher.FileWatcher(NewWatcher, filepath, conn)

	// Wait indefinitely
	select {}
}

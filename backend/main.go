package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Nandgopal-R/LinkFLow/internal/util"
	"github.com/Nandgopal-R/LinkFLow/internal/watcher"
	"github.com/jackc/pgx/v5"
)

func main() {
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

}

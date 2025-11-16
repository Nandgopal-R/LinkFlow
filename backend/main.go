package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
//
// 	"github.com/Nandgopal-R/LinkFLow/internal/util"
// 	"github.com/Nandgopal-R/LinkFLow/internal/watcher"
// 	"github.com/fsnotify/fsnotify"
// 	"github.com/jackc/pgx/v5"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
// 		os.Exit(1)
// 	}
// 	defer conn.Close(context.Background())
//
// 	filepath, err := util.EnsureBlogFileExists()
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Error ensuring blog file exists: %v\n", err)
// 		os.Exit(1)
// 	}
//
// 	NewWatcher, err := fsnotify.NewWatcher() // Create a new file watcher
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer NewWatcher.Close()
//
// 	err = NewWatcher.Add(filepath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	go watcher.FileWatcher(NewWatcher, filepath, conn)
//
// 	<-make(chan bool)
//
// 	// Wait indefinitely
// 	// select {}
// }

import (
	"context"
	"log"
	"net/http"
	"os"

	blogsApi "github.com/Nandgopal-R/LinkFLow/api/blogs"
	"github.com/Nandgopal-R/LinkFLow/cmd"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func SetupRouter() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cmd.DBPool, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Unable to connect to Database")
	}
	defer cmd.DBPool.Close()

	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running",
		})
	})

	lfrouter := r.Group("/linkflow")

	blogsApi.BlogsRoutes(lfrouter)

	err = r.Run()
	if err != nil {
		log.Fatal("Failed to start server")
	}

}

func main() {
	SetupRouter()
}

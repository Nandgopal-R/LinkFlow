package api

import (
	"context"
	"net/http"
	"time"

	"github.com/Nandgopal-R/LinkFLow/cmd"
	db "github.com/Nandgopal-R/LinkFLow/db/gen"

	"github.com/gin-gonic/gin"
)

func FetchBlogs(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	conn, err := cmd.DBPool.Acquire(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to acquire database connection"})
		return
	}

	defer conn.Release()

	q := db.New(conn)

	blogs, err := q.ListBlogsQuery(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch blogs"})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Blogs fetched successfully",
			"blogs":   blogs,
		},
	)

}

func AddBlog(c *gin.Context) {

}

func DeleteBlogById(c *gin.Context) {

}

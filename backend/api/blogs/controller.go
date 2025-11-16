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
	var req db.InsertBlogQueryParams
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	ctx, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()

	tx, err := cmd.DBPool.Begin(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to start transaction"})
	}
	defer tx.Rollback(ctx)

	q := db.New(tx)

	err = q.InsertBlogQuery(ctx, db.InsertBlogQueryParams{
		Title:       req.Title,
		BlogUrl:     req.BlogUrl,
		Description: req.Description,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add blog"})
	}

	if err = tx.Commit(ctx); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to commit transaction"})
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Blog added successfully",
		},
	)
}

func DeleteBlogById(c *gin.Context) {

}

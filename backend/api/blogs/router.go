package api

import "github.com/gin-gonic/gin"

func BlogsRoutes(r *gin.RouterGroup) {
	r.GET("/blogs", FetchBlogs)
	r.POST("/blog", AddBlog)
	r.DELETE("/blogs/:id", DeleteBlogById)
}

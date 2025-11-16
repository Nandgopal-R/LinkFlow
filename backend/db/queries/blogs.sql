-- name: ListBlogsQuery :many
SELECT * FROM blogs;

-- name: InsertBlogQuery :exec
INSERT INTO blogs (title, blog_url, description) 
VALUES($1,$2,$3);

-- name: DeleteBlogQuery :exec
DELETE FROM blogs 
where blog_url=$1;

-- +goose up

-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS "public";
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS blogs (
  id uuid DEFAULT gen_random_uuid() NOT NULL,
  title VARCHAR(255) NOT NULL,
  blog_url TEXT NOT NULL UNIQUE,
  description TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),

  CONSTRAINT blogs_pkey PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tags (
  id uuid DEFAULT gen_random_uuid() NOT NULL,
  name VARCHAR(100) NOT NULL,
  abbreviation VARCHAR(10) NOT NULL,
  CONSTRAINT tags_pkey PRIMARY KEY (id),
)
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS blog_to_tag_mapping (
  id uuid DEFAULT gen_random_uuid() NOT NULL,
  blog_id uuid NOT NULL,
  tag_id uuid NOT NULL,
  CONSTRAINT blog_to_tag_mapping_pkey PRIMARY KEY (id),
  CONSTRAINT fk_blog FOREIGN KEY (blog_id) REFERENCES blogs(id) ON DELETE CASCADE,
  CONSTRAINT fk_tag FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose down
DROP TABLE IF EXISTS blogs;
DROP TABLE IF EXISTS tags;
DROP TABLE IF EXISTS blog_to_tag_mapping;

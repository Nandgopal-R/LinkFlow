-- +goose up

-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS "public";
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS blogs (
  id SERIAL NOT NULL,
  title VARCHAR(255) NOT NULL,
  blog_url TEXT NOT NULL UNIQUE,
  description TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),

  CONSTRAINT blogs_pkey PRIMARY KEY (id)
);
-- +goose StatementEnd
--
-- +goose down
DROP TABLE IF EXISTS blogs;

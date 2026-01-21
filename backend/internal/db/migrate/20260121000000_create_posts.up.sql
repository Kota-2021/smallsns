CREATE TABLE posts (
  post_id     UUID PRIMARY KEY,
  content     TEXT NOT NULL,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE INDEX idx_posts_created_at ON posts (created_at);
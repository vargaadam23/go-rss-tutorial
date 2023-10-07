-- name: CreatePost :exec
-- type: CreatePost
INSERT INTO posts (
  created_at, updated_at, url, feed_id, title, description, published_at
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
);

-- name: GetPostsForFeed :many
SELECT * FROM posts WHERE feed_id=?;

-- name: GetPostsForUser :many
SELECT posts.* FROM posts 
JOIN feed_follows ON posts.feed_id=feed_follows.feed_id
WHERE feed_follows.user_id=?
ORDER BY posts.published_at LIMIT ?;

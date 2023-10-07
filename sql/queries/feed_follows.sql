-- name: CreateFeedFollow :exec
-- type: CreateFeedFollow
INSERT INTO feed_follows (
  created_at, updated_at, user_id, feed_id
) VALUES (
  ?, ?, ?, ?
);

-- name: GetFeedFollowByUser :many
SELECT * FROM feed_follows WHERE user_id=?;

-- name: GetFeedFollowByFeed :many
SELECT * FROM feed_follows WHERE feed_id=?;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id=? AND user_id=?;
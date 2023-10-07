-- name: CreateFeed :exec
-- type: CreateFeed
INSERT INTO feeds (
  created_at, updated_at, name, url, user_id
) VALUES (
  ?, ?, ?, ?, ?
);

-- name: GetFeedsForUser :many
SELECT * FROM feeds WHERE user_id=?;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds ORDER BY last_fetched_at ASC LIMIT ?;

-- name: MarkFeedFetched :exec
UPDATE feeds SET last_fetched_at=NOW(), updated_at=NOW() WHERE id=?;

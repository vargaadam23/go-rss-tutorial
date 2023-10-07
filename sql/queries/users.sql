-- name: CreateUser :exec
-- type: CreateUser
INSERT INTO users (
  created_at, updated_at, name, api_key
) VALUES (
  ?, ?, ?, ?
);

-- name: GetUserByApiKey :one
SELECT * FROM users WHERE api_key = ?
// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: users.sql

package database

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (
  created_at, updated_at, name, api_key
) VALUES (
  ?, ?, ?, ?
)
`

type CreateUserParams struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}

// type: CreateUser
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.ExecContext(ctx, createUser,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Name,
		arg.ApiKey,
	)
	return err
}

const getUserByApiKey = `-- name: GetUserByApiKey :one
SELECT id, created_at, updated_at, name, api_key FROM users WHERE api_key = ?
`

func (q *Queries) GetUserByApiKey(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByApiKey, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.ApiKey,
	)
	return i, err
}

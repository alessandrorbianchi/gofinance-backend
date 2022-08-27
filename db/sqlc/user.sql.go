// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  st_username,
  st_password,
  st_email
) VALUES (
  $1, $2, $3
) RETURNING id, st_username, st_password, st_email, dt_created_at
`

type CreateUserParams struct {
	StUsername string `json:"st_username"`
	StPassword string `json:"st_password"`
	StEmail    string `json:"st_email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.StUsername, arg.StPassword, arg.StEmail)
	var i User
	err := row.Scan(
		&i.ID,
		&i.StUsername,
		&i.StPassword,
		&i.StEmail,
		&i.DtCreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, st_username, st_password, st_email, dt_created_at FROM users
WHERE st_username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, stUsername string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, stUsername)
	var i User
	err := row.Scan(
		&i.ID,
		&i.StUsername,
		&i.StPassword,
		&i.StEmail,
		&i.DtCreatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, st_username, st_password, st_email, dt_created_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.StUsername,
		&i.StPassword,
		&i.StEmail,
		&i.DtCreatedAt,
	)
	return i, err
}

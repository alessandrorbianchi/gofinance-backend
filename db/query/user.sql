-- name: CreateUser :one
INSERT INTO users (
  st_username,
  st_password,
  st_email
) VALUES (
  $1, $2, $3
) RETURNING *;
    
-- name: GetUser :one
SELECT * FROM users
WHERE st_username = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;
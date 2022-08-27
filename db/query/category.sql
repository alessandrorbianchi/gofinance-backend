-- name: CreateCategory :one
INSERT INTO categories (
  co_user_id,
  st_title,
  st_type,
  st_description
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1;

-- name: GetCategories :many
SELECT * FROM categories 
where co_user_id = $1 and st_type = $2 
and st_title like $3 and st_description like $4;

-- name: GetCategoriesByUserIdAndType :many
SELECT * FROM categories 
where co_user_id = $1 and st_type = $2;

-- name: GetCategoriesByUserIdAndTypeAndTitle :many
SELECT * FROM categories 
where co_user_id = $1 and st_type = $2
and st_title like $3;

-- name: GetCategoriesByUserIdAndTypeAndDescription :many
SELECT * FROM categories 
where co_user_id = $1 and st_type = $2
and st_description like $3;

-- name: UpdateCategories :one
UPDATE categories
SET st_title = $2, st_description = $3
WHERE id = $1
RETURNING *;

-- name: DeleteCategories :exec
DELETE FROM categories
WHERE id = $1;
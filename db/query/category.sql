-- name: CreateCategory :one
INSERT INTO categories (
  user_id,
  title,
  type,
  description
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories 
WHERE id = $1 LIMIT 1;

-- name: GetCategories :many
SELECT * FROM categories 
where user_id = $1 and type = $2 
and lower(title) like CONCAT('%', lower(@title::text), '%') 
and lower(description) like CONCAT('%', lower(@description::text), '%');

-- name: UpdateCategories :one
UPDATE categories 
SET title = $2, description = $3 
WHERE id = $1 RETURNING *;

-- name: DeleteCategories :exec
DELETE FROM categories 
WHERE id = $1;
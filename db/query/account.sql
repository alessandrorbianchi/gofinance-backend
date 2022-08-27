-- name: CreateAccount :one
INSERT INTO accounts (
  co_user_id,
  co_category_id,
  st_title,
  st_type,
  st_description,
  vl_value,
  dt_date
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccounts :many
SELECT 
  a.id,
  a.co_user_id,
  a.st_title,
  a.st_type,
  a.st_description,
  a.vl_value,
  a.dt_date,
  a.dt_created_at,
  c.st_title as category_st_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.co_category_id 
where a.co_user_id = $1 and a.st_type = $2 
and a.co_category_id = $3 and a.st_title like $4
and a.st_description like $5 and a.dt_date = $6;

-- name: GetAccountsByUserIdAndst_type :many
SELECT 
  a.id,
  a.co_user_id,
  a.st_title,
  a.st_type,
  a.st_description,
  a.vl_value,
  a.dt_date,
  a.dt_created_at,
  c.st_title as category_st_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.co_category_id 
where a.co_user_id = $1 and a.st_type = $2;

-- name: GetAccountsByUserIdAndst_typeAndCategoryId :many
SELECT 
  a.id,
  a.co_user_id,
  a.st_title,
  a.st_type,
  a.st_description,
  a.vl_value,
  a.dt_date,
  a.dt_created_at,
  c.st_title as category_st_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.co_category_id 
where a.co_user_id = $1 and a.st_type = $2 
and a.co_category_id = $3;

-- name: GetAccountsByUserIdAndst_typeAndCategoryIdAndst_title :many
SELECT 
  a.id,
  a.co_user_id,
  a.st_title,
  a.st_type,
  a.st_description,
  a.vl_value,
  a.dt_date,
  a.dt_created_at,
  c.st_title as category_st_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.co_category_id 
where a.co_user_id = $1 and a.st_type = $2 
and a.co_category_id = $3 and a.st_title like $4;

-- name: GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleAndst_description :many
SELECT 
  a.id,
  a.co_user_id,
  a.st_title,
  a.st_type,
  a.st_description,
  a.vl_value,
  a.dt_date,
  a.dt_created_at,
  c.st_title as category_st_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.co_category_id 
where a.co_user_id = $1 and a.st_type = $2 
and a.co_category_id = $3 and a.st_title like $4
and a.st_description like $5;

-- name: GetAccountsByUserIdAndst_typeAndst_title :many
SELECT 
  a.id,
  a.co_user_id,
  a.st_title,
  a.st_type,
  a.st_description,
  a.vl_value,
  a.dt_date,
  a.dt_created_at,
  c.st_title as category_st_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.co_category_id 
where a.co_user_id = $1 and a.st_type = $2
and a.st_title like $3;

-- name: GetAccountsByUserIdAndst_typeAndst_description :many
SELECT 
  a.id,
  a.co_user_id,
  a.st_title,
  a.st_type,
  a.st_description,
  a.vl_value,
  a.dt_date,
  a.dt_created_at,
  c.st_title as category_st_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.co_category_id 
where a.co_user_id = $1 and a.st_type = $2
and a.st_description like $3;

-- name: GetAccountsByUserIdAndst_typeAnddt_date :many
SELECT 
  a.id,
  a.co_user_id,
  a.st_title,
  a.st_type,
  a.st_description,
  a.vl_value,
  a.dt_date,
  a.dt_created_at,
  c.st_title as category_st_title
FROM accounts a
LEFT JOIN categories c ON c.id = a.co_category_id 
where a.co_user_id = $1 and a.st_type = $2
and a.dt_date like $3;

-- name: GetAccountsReports :one
SELECT SUM(vl_value) AS sum_value FROM accounts 
where co_user_id = $1 and st_type = $2;

-- name: GetAccountsGraph :one
SELECT COUNT(*) FROM accounts 
where co_user_id = $1 and st_type = $2;

-- name: UpdateAccount :one
UPDATE accounts
SET st_title = $2, st_description = $3, vl_value = $4
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1;
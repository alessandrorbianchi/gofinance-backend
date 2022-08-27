// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: account.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createAccount = `-- name: CreateAccount :one
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
) RETURNING id, co_user_id, co_category_id, st_title, st_type, st_description, vl_value, dt_date, dt_created_at
`

type CreateAccountParams struct {
	CoUserID      int32     `json:"co_user_id"`
	CoCategoryID  int32     `json:"co_category_id"`
	StTitle       string    `json:"st_title"`
	StType        string    `json:"st_type"`
	StDescription string    `json:"st_description"`
	VlValue       int32     `json:"vl_value"`
	DtDate        time.Time `json:"dt_date"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.CoUserID,
		arg.CoCategoryID,
		arg.StTitle,
		arg.StType,
		arg.StDescription,
		arg.VlValue,
		arg.DtDate,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.CoUserID,
		&i.CoCategoryID,
		&i.StTitle,
		&i.StType,
		&i.StDescription,
		&i.VlValue,
		&i.DtDate,
		&i.DtCreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, co_user_id, co_category_id, st_title, st_type, st_description, vl_value, dt_date, dt_created_at FROM accounts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int32) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.CoUserID,
		&i.CoCategoryID,
		&i.StTitle,
		&i.StType,
		&i.StDescription,
		&i.VlValue,
		&i.DtDate,
		&i.DtCreatedAt,
	)
	return i, err
}

const getAccounts = `-- name: GetAccounts :many
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
and a.st_description like $5 and a.dt_date = $6
`

type GetAccountsParams struct {
	CoUserID      int32     `json:"co_user_id"`
	StType        string    `json:"st_type"`
	CoCategoryID  int32     `json:"co_category_id"`
	StTitle       string    `json:"st_title"`
	StDescription string    `json:"st_description"`
	DtDate        time.Time `json:"dt_date"`
}

type GetAccountsRow struct {
	ID              int32          `json:"id"`
	CoUserID        int32          `json:"co_user_id"`
	StTitle         string         `json:"st_title"`
	StType          string         `json:"st_type"`
	StDescription   string         `json:"st_description"`
	VlValue         int32          `json:"vl_value"`
	DtDate          time.Time      `json:"dt_date"`
	DtCreatedAt     time.Time      `json:"dt_created_at"`
	CategoryStTitle sql.NullString `json:"category_st_title"`
}

func (q *Queries) GetAccounts(ctx context.Context, arg GetAccountsParams) ([]GetAccountsRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccounts,
		arg.CoUserID,
		arg.StType,
		arg.CoCategoryID,
		arg.StTitle,
		arg.StDescription,
		arg.DtDate,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsRow{}
	for rows.Next() {
		var i GetAccountsRow
		if err := rows.Scan(
			&i.ID,
			&i.CoUserID,
			&i.StTitle,
			&i.StType,
			&i.StDescription,
			&i.VlValue,
			&i.DtDate,
			&i.DtCreatedAt,
			&i.CategoryStTitle,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccountsByUserIdAndst_type = `-- name: GetAccountsByUserIdAndst_type :many
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
`

type GetAccountsByUserIdAndst_typeParams struct {
	CoUserID int32  `json:"co_user_id"`
	StType   string `json:"st_type"`
}

type GetAccountsByUserIdAndst_typeRow struct {
	ID              int32          `json:"id"`
	CoUserID        int32          `json:"co_user_id"`
	StTitle         string         `json:"st_title"`
	StType          string         `json:"st_type"`
	StDescription   string         `json:"st_description"`
	VlValue         int32          `json:"vl_value"`
	DtDate          time.Time      `json:"dt_date"`
	DtCreatedAt     time.Time      `json:"dt_created_at"`
	CategoryStTitle sql.NullString `json:"category_st_title"`
}

func (q *Queries) GetAccountsByUserIdAndst_type(ctx context.Context, arg GetAccountsByUserIdAndst_typeParams) ([]GetAccountsByUserIdAndst_typeRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUserIdAndst_type, arg.CoUserID, arg.StType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsByUserIdAndst_typeRow{}
	for rows.Next() {
		var i GetAccountsByUserIdAndst_typeRow
		if err := rows.Scan(
			&i.ID,
			&i.CoUserID,
			&i.StTitle,
			&i.StType,
			&i.StDescription,
			&i.VlValue,
			&i.DtDate,
			&i.DtCreatedAt,
			&i.CategoryStTitle,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccountsByUserIdAndst_typeAndCategoryId = `-- name: GetAccountsByUserIdAndst_typeAndCategoryId :many
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
and a.co_category_id = $3
`

type GetAccountsByUserIdAndst_typeAndCategoryIdParams struct {
	CoUserID     int32  `json:"co_user_id"`
	StType       string `json:"st_type"`
	CoCategoryID int32  `json:"co_category_id"`
}

type GetAccountsByUserIdAndst_typeAndCategoryIdRow struct {
	ID              int32          `json:"id"`
	CoUserID        int32          `json:"co_user_id"`
	StTitle         string         `json:"st_title"`
	StType          string         `json:"st_type"`
	StDescription   string         `json:"st_description"`
	VlValue         int32          `json:"vl_value"`
	DtDate          time.Time      `json:"dt_date"`
	DtCreatedAt     time.Time      `json:"dt_created_at"`
	CategoryStTitle sql.NullString `json:"category_st_title"`
}

func (q *Queries) GetAccountsByUserIdAndst_typeAndCategoryId(ctx context.Context, arg GetAccountsByUserIdAndst_typeAndCategoryIdParams) ([]GetAccountsByUserIdAndst_typeAndCategoryIdRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUserIdAndst_typeAndCategoryId, arg.CoUserID, arg.StType, arg.CoCategoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsByUserIdAndst_typeAndCategoryIdRow{}
	for rows.Next() {
		var i GetAccountsByUserIdAndst_typeAndCategoryIdRow
		if err := rows.Scan(
			&i.ID,
			&i.CoUserID,
			&i.StTitle,
			&i.StType,
			&i.StDescription,
			&i.VlValue,
			&i.DtDate,
			&i.DtCreatedAt,
			&i.CategoryStTitle,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccountsByUserIdAndst_typeAndCategoryIdAndst_title = `-- name: GetAccountsByUserIdAndst_typeAndCategoryIdAndst_title :many
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
`

type GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleParams struct {
	CoUserID     int32  `json:"co_user_id"`
	StType       string `json:"st_type"`
	CoCategoryID int32  `json:"co_category_id"`
	StTitle      string `json:"st_title"`
}

type GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleRow struct {
	ID              int32          `json:"id"`
	CoUserID        int32          `json:"co_user_id"`
	StTitle         string         `json:"st_title"`
	StType          string         `json:"st_type"`
	StDescription   string         `json:"st_description"`
	VlValue         int32          `json:"vl_value"`
	DtDate          time.Time      `json:"dt_date"`
	DtCreatedAt     time.Time      `json:"dt_created_at"`
	CategoryStTitle sql.NullString `json:"category_st_title"`
}

func (q *Queries) GetAccountsByUserIdAndst_typeAndCategoryIdAndst_title(ctx context.Context, arg GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleParams) ([]GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUserIdAndst_typeAndCategoryIdAndst_title,
		arg.CoUserID,
		arg.StType,
		arg.CoCategoryID,
		arg.StTitle,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleRow{}
	for rows.Next() {
		var i GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleRow
		if err := rows.Scan(
			&i.ID,
			&i.CoUserID,
			&i.StTitle,
			&i.StType,
			&i.StDescription,
			&i.VlValue,
			&i.DtDate,
			&i.DtCreatedAt,
			&i.CategoryStTitle,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccountsByUserIdAndst_typeAndCategoryIdAndst_titleAndst_description = `-- name: GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleAndst_description :many
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
and a.st_description like $5
`

type GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleAndst_descriptionParams struct {
	CoUserID      int32  `json:"co_user_id"`
	StType        string `json:"st_type"`
	CoCategoryID  int32  `json:"co_category_id"`
	StTitle       string `json:"st_title"`
	StDescription string `json:"st_description"`
}

type GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleAndst_descriptionRow struct {
	ID              int32          `json:"id"`
	CoUserID        int32          `json:"co_user_id"`
	StTitle         string         `json:"st_title"`
	StType          string         `json:"st_type"`
	StDescription   string         `json:"st_description"`
	VlValue         int32          `json:"vl_value"`
	DtDate          time.Time      `json:"dt_date"`
	DtCreatedAt     time.Time      `json:"dt_created_at"`
	CategoryStTitle sql.NullString `json:"category_st_title"`
}

func (q *Queries) GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleAndst_description(ctx context.Context, arg GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleAndst_descriptionParams) ([]GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleAndst_descriptionRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUserIdAndst_typeAndCategoryIdAndst_titleAndst_description,
		arg.CoUserID,
		arg.StType,
		arg.CoCategoryID,
		arg.StTitle,
		arg.StDescription,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleAndst_descriptionRow{}
	for rows.Next() {
		var i GetAccountsByUserIdAndst_typeAndCategoryIdAndst_titleAndst_descriptionRow
		if err := rows.Scan(
			&i.ID,
			&i.CoUserID,
			&i.StTitle,
			&i.StType,
			&i.StDescription,
			&i.VlValue,
			&i.DtDate,
			&i.DtCreatedAt,
			&i.CategoryStTitle,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccountsByUserIdAndst_typeAnddt_date = `-- name: GetAccountsByUserIdAndst_typeAnddt_date :many
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
and a.dt_date like $3
`

type GetAccountsByUserIdAndst_typeAnddt_dateParams struct {
	CoUserID int32     `json:"co_user_id"`
	StType   string    `json:"st_type"`
	DtDate   time.Time `json:"dt_date"`
}

type GetAccountsByUserIdAndst_typeAnddt_dateRow struct {
	ID              int32          `json:"id"`
	CoUserID        int32          `json:"co_user_id"`
	StTitle         string         `json:"st_title"`
	StType          string         `json:"st_type"`
	StDescription   string         `json:"st_description"`
	VlValue         int32          `json:"vl_value"`
	DtDate          time.Time      `json:"dt_date"`
	DtCreatedAt     time.Time      `json:"dt_created_at"`
	CategoryStTitle sql.NullString `json:"category_st_title"`
}

func (q *Queries) GetAccountsByUserIdAndst_typeAnddt_date(ctx context.Context, arg GetAccountsByUserIdAndst_typeAnddt_dateParams) ([]GetAccountsByUserIdAndst_typeAnddt_dateRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUserIdAndst_typeAnddt_date, arg.CoUserID, arg.StType, arg.DtDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsByUserIdAndst_typeAnddt_dateRow{}
	for rows.Next() {
		var i GetAccountsByUserIdAndst_typeAnddt_dateRow
		if err := rows.Scan(
			&i.ID,
			&i.CoUserID,
			&i.StTitle,
			&i.StType,
			&i.StDescription,
			&i.VlValue,
			&i.DtDate,
			&i.DtCreatedAt,
			&i.CategoryStTitle,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccountsByUserIdAndst_typeAndst_description = `-- name: GetAccountsByUserIdAndst_typeAndst_description :many
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
and a.st_description like $3
`

type GetAccountsByUserIdAndst_typeAndst_descriptionParams struct {
	CoUserID      int32  `json:"co_user_id"`
	StType        string `json:"st_type"`
	StDescription string `json:"st_description"`
}

type GetAccountsByUserIdAndst_typeAndst_descriptionRow struct {
	ID              int32          `json:"id"`
	CoUserID        int32          `json:"co_user_id"`
	StTitle         string         `json:"st_title"`
	StType          string         `json:"st_type"`
	StDescription   string         `json:"st_description"`
	VlValue         int32          `json:"vl_value"`
	DtDate          time.Time      `json:"dt_date"`
	DtCreatedAt     time.Time      `json:"dt_created_at"`
	CategoryStTitle sql.NullString `json:"category_st_title"`
}

func (q *Queries) GetAccountsByUserIdAndst_typeAndst_description(ctx context.Context, arg GetAccountsByUserIdAndst_typeAndst_descriptionParams) ([]GetAccountsByUserIdAndst_typeAndst_descriptionRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUserIdAndst_typeAndst_description, arg.CoUserID, arg.StType, arg.StDescription)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsByUserIdAndst_typeAndst_descriptionRow{}
	for rows.Next() {
		var i GetAccountsByUserIdAndst_typeAndst_descriptionRow
		if err := rows.Scan(
			&i.ID,
			&i.CoUserID,
			&i.StTitle,
			&i.StType,
			&i.StDescription,
			&i.VlValue,
			&i.DtDate,
			&i.DtCreatedAt,
			&i.CategoryStTitle,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccountsByUserIdAndst_typeAndst_title = `-- name: GetAccountsByUserIdAndst_typeAndst_title :many
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
and a.st_title like $3
`

type GetAccountsByUserIdAndst_typeAndst_titleParams struct {
	CoUserID int32  `json:"co_user_id"`
	StType   string `json:"st_type"`
	StTitle  string `json:"st_title"`
}

type GetAccountsByUserIdAndst_typeAndst_titleRow struct {
	ID              int32          `json:"id"`
	CoUserID        int32          `json:"co_user_id"`
	StTitle         string         `json:"st_title"`
	StType          string         `json:"st_type"`
	StDescription   string         `json:"st_description"`
	VlValue         int32          `json:"vl_value"`
	DtDate          time.Time      `json:"dt_date"`
	DtCreatedAt     time.Time      `json:"dt_created_at"`
	CategoryStTitle sql.NullString `json:"category_st_title"`
}

func (q *Queries) GetAccountsByUserIdAndst_typeAndst_title(ctx context.Context, arg GetAccountsByUserIdAndst_typeAndst_titleParams) ([]GetAccountsByUserIdAndst_typeAndst_titleRow, error) {
	rows, err := q.db.QueryContext(ctx, getAccountsByUserIdAndst_typeAndst_title, arg.CoUserID, arg.StType, arg.StTitle)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetAccountsByUserIdAndst_typeAndst_titleRow{}
	for rows.Next() {
		var i GetAccountsByUserIdAndst_typeAndst_titleRow
		if err := rows.Scan(
			&i.ID,
			&i.CoUserID,
			&i.StTitle,
			&i.StType,
			&i.StDescription,
			&i.VlValue,
			&i.DtDate,
			&i.DtCreatedAt,
			&i.CategoryStTitle,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getAccountsGraph = `-- name: GetAccountsGraph :one
SELECT COUNT(*) FROM accounts 
where co_user_id = $1 and st_type = $2
`

type GetAccountsGraphParams struct {
	CoUserID int32  `json:"co_user_id"`
	StType   string `json:"st_type"`
}

func (q *Queries) GetAccountsGraph(ctx context.Context, arg GetAccountsGraphParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAccountsGraph, arg.CoUserID, arg.StType)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getAccountsReports = `-- name: GetAccountsReports :one
SELECT SUM(vl_value) AS sum_value FROM accounts 
where co_user_id = $1 and st_type = $2
`

type GetAccountsReportsParams struct {
	CoUserID int32  `json:"co_user_id"`
	StType   string `json:"st_type"`
}

func (q *Queries) GetAccountsReports(ctx context.Context, arg GetAccountsReportsParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, getAccountsReports, arg.CoUserID, arg.StType)
	var sum_value int64
	err := row.Scan(&sum_value)
	return sum_value, err
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts
SET st_title = $2, st_description = $3, vl_value = $4
WHERE id = $1
RETURNING id, co_user_id, co_category_id, st_title, st_type, st_description, vl_value, dt_date, dt_created_at
`

type UpdateAccountParams struct {
	ID            int32  `json:"id"`
	StTitle       string `json:"st_title"`
	StDescription string `json:"st_description"`
	VlValue       int32  `json:"vl_value"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount,
		arg.ID,
		arg.StTitle,
		arg.StDescription,
		arg.VlValue,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.CoUserID,
		&i.CoCategoryID,
		&i.StTitle,
		&i.StType,
		&i.StDescription,
		&i.VlValue,
		&i.DtDate,
		&i.DtCreatedAt,
	)
	return i, err
}
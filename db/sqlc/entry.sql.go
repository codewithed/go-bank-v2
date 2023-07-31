// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: entry.sql

package db

import (
	"context"
)

const createEntry = `-- name: CreateEntry :one
INSERT INTO entries (
    amount, account_id
)
VALUES ($1, $2)
RETURNING id, amount, account_id, created_at
`

type CreateEntryParams struct {
	Amount    int64 `json:"amount"`
	AccountID int64 `json:"account_id"`
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.Amount, arg.AccountID)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.AccountID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteEntry = `-- name: DeleteEntry :exec
DELETE FROM entries WHERE id = $1
`

func (q *Queries) DeleteEntry(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteEntry, id)
	return err
}

const getEntry = `-- name: GetEntry :one
SELECT id, amount, account_id, created_at FROM entries
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetEntry(ctx context.Context, id int64) (Entry, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.Amount,
		&i.AccountID,
		&i.CreatedAt,
	)
	return i, err
}

const listEntries = `-- name: ListEntries :many
SELECT id, amount, account_id, created_at FROM entries
LIMIT $1
OFFSET $2
`

type ListEntriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error) {
	rows, err := q.db.QueryContext(ctx, listEntries, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Entry{}
	for rows.Next() {
		var i Entry
		if err := rows.Scan(
			&i.ID,
			&i.Amount,
			&i.AccountID,
			&i.CreatedAt,
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

const updateEntry = `-- name: UpdateEntry :exec
UPDATE entries 
SET amount = $2, account_id = $3
WHERE id = $1
`

type UpdateEntryParams struct {
	ID        int64 `json:"id"`
	Amount    int64 `json:"amount"`
	AccountID int64 `json:"account_id"`
}

func (q *Queries) UpdateEntry(ctx context.Context, arg UpdateEntryParams) error {
	_, err := q.db.ExecContext(ctx, updateEntry, arg.ID, arg.Amount, arg.AccountID)
	return err
}
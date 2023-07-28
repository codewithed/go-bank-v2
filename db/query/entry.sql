-- name: CreateEntry :one
INSERT INTO entries (
    amount, account_id
)
VALUES ($1, $2)
RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntries :many
SELECT * FROM entries
LIMIT $1
OFFSET $2;

-- name: UpdateEntry :exec
UPDATE entries 
SET amount = $2, account_id = $3
WHERE id = $1;

-- name: DeleteEntry :exec
DELETE FROM entries WHERE id = $1;
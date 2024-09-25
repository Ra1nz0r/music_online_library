-- name: Add :one
INSERT INTO library ("group", song)
VALUES ($1, $2)
RETURNING *;
-- name: Delete :exec
DELETE FROM library
WHERE id = $1;
-- name: FetchParam :exec
UPDATE library
SET "releaseDate" = $2,
    text = $3,
    patronymic = $4
WHERE id = $1;
-- name: GetOne :one
SELECT *
FROM library
WHERE id = $1
LIMIT 1;
-- name: ListAll :many
SELECT *
FROM library
ORDER BY id;
-- name: List :many
SELECT *
FROM library
ORDER BY id
LIMIT $1 OFFSET $2;
-- name: Update :exec
UPDATE library
SET "group" = $2,
    song = $3,
    "releaseDate" = $4,
    text = $5,
    patronymic = $6
WHERE id = $1;
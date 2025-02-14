-- name: CreateFeed :one

INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;


-- name: GetFeeds :many
SELECT feeds.name AS name, feeds.url AS url, users.name AS users_name FROM feeds JOIN users ON feeds.user_id = users.id; 


-- name: GetFeed :one
SELECT id, name, url FROM feeds WHERE url=$1; 


-- name: DeleteFeed :exec
DELETE FROM feeds where feeds.url=$1;  

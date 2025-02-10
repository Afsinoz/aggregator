-- name: CreateFeedFollow :one 
WITH inserted_feed_follow AS (
INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING *)
SELECT inserted_feed_follow.*, feeds.name, users.name FROM inserted_feed_follow 
INNER JOIN users ON user_id = users.id
INNER JOIN feeds ON feed_id = feeds.id; 


-- name: GetFeedFollowsForUsers :many 
SELECT *, feeds.name as feed_name, users.name as user_name FROM feed_follows
INNER JOIN users ON user_id = users.id
INNER JOIN feeds ON feed_id = feeds.id
WHERE $1 = users.name;



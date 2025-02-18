-- name: MarkFeedFetched :exec

UPDATE feeds 
SET last_fetched_at = NOW(), updated_at = NOW(); 

-- name: GetNextFeedToFetch :one

SELECT * FROM feeds 
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1; 

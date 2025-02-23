// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: feed_follows.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
INSERT INTO feed_follows(id, created_at, updated_at, user_id, feed_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) RETURNING id, created_at, updated_at, user_id, feed_id)
SELECT inserted_feed_follow.id, inserted_feed_follow.created_at, inserted_feed_follow.updated_at, inserted_feed_follow.user_id, inserted_feed_follow.feed_id, feeds.name, users.name FROM inserted_feed_follow 
INNER JOIN users ON user_id = users.id
INNER JOIN feeds ON feed_id = feeds.id
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
}

type CreateFeedFollowRow struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    uuid.UUID
	FeedID    uuid.UUID
	Name      string
	Name_2    string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
		arg.FeedID,
	)
	var i CreateFeedFollowRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.FeedID,
		&i.Name,
		&i.Name_2,
	)
	return i, err
}

const getFeedFollowsForUsers = `-- name: GetFeedFollowsForUsers :many
SELECT feed_follows.id, feed_follows.created_at, feed_follows.updated_at, feed_follows.user_id, feed_id, users.id, users.created_at, users.updated_at, users.name, feeds.id, feeds.created_at, feeds.updated_at, feeds.name, url, feeds.user_id, last_fetched_at, feeds.name as feed_name, users.name as user_name FROM feed_follows
INNER JOIN users ON user_id = users.id
INNER JOIN feeds ON feed_id = feeds.id
WHERE $1 = users.name
`

type GetFeedFollowsForUsersRow struct {
	ID            uuid.UUID
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserID        uuid.UUID
	FeedID        uuid.UUID
	ID_2          uuid.UUID
	CreatedAt_2   time.Time
	UpdatedAt_2   time.Time
	Name          string
	ID_3          uuid.UUID
	CreatedAt_3   time.Time
	UpdatedAt_3   time.Time
	Name_2        string
	Url           string
	UserID_2      uuid.UUID
	LastFetchedAt sql.NullTime
	FeedName      string
	UserName      string
}

func (q *Queries) GetFeedFollowsForUsers(ctx context.Context, name string) ([]GetFeedFollowsForUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsForUsers, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedFollowsForUsersRow
	for rows.Next() {
		var i GetFeedFollowsForUsersRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.FeedID,
			&i.ID_2,
			&i.CreatedAt_2,
			&i.UpdatedAt_2,
			&i.Name,
			&i.ID_3,
			&i.CreatedAt_3,
			&i.UpdatedAt_3,
			&i.Name_2,
			&i.Url,
			&i.UserID_2,
			&i.LastFetchedAt,
			&i.FeedName,
			&i.UserName,
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

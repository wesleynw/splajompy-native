// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: notifications.sql

package db

import (
	"context"
)

const getNotificationById = `-- name: GetNotificationById :one
SELECT notification_id, user_id, post_id, comment_id, message, link, viewed, created_at
FROM notifications
WHERE notification_id = $1
LIMIT 1
`

func (q *Queries) GetNotificationById(ctx context.Context, notificationID int32) (Notification, error) {
	row := q.db.QueryRow(ctx, getNotificationById, notificationID)
	var i Notification
	err := row.Scan(
		&i.NotificationID,
		&i.UserID,
		&i.PostID,
		&i.CommentID,
		&i.Message,
		&i.Link,
		&i.Viewed,
		&i.CreatedAt,
	)
	return i, err
}

const getNotificationsForUserId = `-- name: GetNotificationsForUserId :many
SELECT notification_id, user_id, post_id, comment_id, message, link, viewed, created_at
FROM notifications 
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3
`

type GetNotificationsForUserIdParams struct {
	UserID int32 `json:"userId"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetNotificationsForUserId(ctx context.Context, arg GetNotificationsForUserIdParams) ([]Notification, error) {
	rows, err := q.db.Query(ctx, getNotificationsForUserId, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Notification
	for rows.Next() {
		var i Notification
		if err := rows.Scan(
			&i.NotificationID,
			&i.UserID,
			&i.PostID,
			&i.CommentID,
			&i.Message,
			&i.Link,
			&i.Viewed,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const markAllNotificationsAsReadForUser = `-- name: MarkAllNotificationsAsReadForUser :exec
UPDATE notifications
SET viewed = TRUE
WHERE user_id = $1
`

func (q *Queries) MarkAllNotificationsAsReadForUser(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, markAllNotificationsAsReadForUser, userID)
	return err
}

const markNotificationAsReadById = `-- name: MarkNotificationAsReadById :exec
UPDATE notifications
SET viewed = TRUE
WHERE notification_id = $1
`

func (q *Queries) MarkNotificationAsReadById(ctx context.Context, notificationID int32) error {
	_, err := q.db.Exec(ctx, markNotificationAsReadById, notificationID)
	return err
}

const userHasUnreadNotifications = `-- name: UserHasUnreadNotifications :one
SELECT EXISTS (
  SELECT 1
  FROM notifications
  WHERE user_id = $1 AND viewed = FALSE
)
`

func (q *Queries) UserHasUnreadNotifications(ctx context.Context, userID int32) (bool, error) {
	row := q.db.QueryRow(ctx, userHasUnreadNotifications, userID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

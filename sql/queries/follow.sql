-- name: CreateFeedFollow :one
INSERT INTO  feed_follow (id,created_at,updated_at , user_id , feed_id)
VALUES($1,$2,$3,$4,$5 )
 RETURNING *;

-- name: GetFeedFollowsByUser :many
SELECT * FROM feed_follow where user_id=$1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follow where id=$1 and user_id=$2;
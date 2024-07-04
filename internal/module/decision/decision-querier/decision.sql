-- name: UpsertOne :one
INSERT INTO decisions (actor_user_id, recipient_user_id, liked_recipient, recipient_likes_actor, created_at)
VALUES ($1, $2, $3, $4, NOW())
ON CONFLICT (actor_user_id, recipient_user_id)
DO UPDATE SET
    liked_recipient = EXCLUDED.liked_recipient,
    recipient_likes_actor = EXCLUDED.recipient_likes_actor,
    created_at = EXCLUDED.created_at
RETURNING *;

-- name: GetOneByActorUserId :one
SELECT * FROM decisions
WHERE actor_user_id = $1;

-- name: GetOneByRecipientUserID :one
SELECT * FROM decisions
WHERE recipient_user_id = $1;

-- name: List :many
SELECT *
FROM decisions
WHERE
  (actor_user_id = sqlc.narg('actor_user_id') OR sqlc.narg('actor_user_id') IS NULL) 
  AND
  (recipient_user_id = sqlc.narg('recipient_user_id') OR sqlc.narg('recipient_user_id') IS NULL) 
  AND
  (liked_recipient = sqlc.narg('liked_recipient') OR sqlc.narg('liked_recipient') IS NULL)
  AND
  (recipient_likes_actor = sqlc.narg('recipient_likes_actor') OR sqlc.narg('recipient_likes_actor') IS NULL)
  AND
  (created_at < sqlc.arg('pagination_token')::date OR sqlc.narg('pagination_token') IS NULL)
ORDER BY created_at DESC
LIMIT sqlc.narg('limit');

-- name: Count :one
SELECT COUNT(*)
FROM decisions
WHERE
  (actor_user_id = sqlc.narg('actor_user_id') OR sqlc.narg('actor_user_id') IS NULL) 
  AND
  (recipient_user_id = sqlc.narg('recipient_user_id') OR sqlc.narg('recipient_user_id') IS NULL) 
  AND
  (liked_recipient = sqlc.narg('liked_recipient') OR sqlc.narg('liked_recipient') IS NULL)
  AND
  (recipient_likes_actor = sqlc.narg('recipient_likes_actor') OR sqlc.narg('recipient_likes_actor') IS NULL);

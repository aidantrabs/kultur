-- name: CreateSubscription :one
INSERT INTO subscriptions (
    email, digest_weekly, festival_reminders, confirmation_token, unsubscribe_token
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetSubscriptionByEmail :one
SELECT * FROM subscriptions
WHERE email = $1;

-- name: GetSubscriptionByConfirmationToken :one
SELECT * FROM subscriptions
WHERE confirmation_token = $1;

-- name: GetSubscriptionByUnsubscribeToken :one
SELECT * FROM subscriptions
WHERE unsubscribe_token = $1;

-- name: ConfirmSubscription :exec
UPDATE subscriptions
SET confirmed = true, confirmation_token = NULL
WHERE id = $1;

-- name: DeleteSubscription :exec
DELETE FROM subscriptions
WHERE id = $1;

-- name: ListConfirmedWeeklyDigest :many
SELECT * FROM subscriptions
WHERE confirmed = true AND digest_weekly = true;

-- name: ListAllSubscriptions :many
SELECT * FROM subscriptions
ORDER BY created_at DESC;

-- name: ListMemoriesByFestival :many
SELECT * FROM memories
WHERE festival_id = $1 AND status = 'approved'
ORDER BY submitted_at DESC;

-- name: CreateMemory :one
INSERT INTO memories (
    festival_id, author_name, author_email, content, year_of_memory
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetMemoryByID :one
SELECT * FROM memories
WHERE id = $1;

-- name: UpdateMemoryStatus :exec
UPDATE memories
SET status = $2
WHERE id = $1;

-- name: ListAllMemories :many
SELECT * FROM memories
ORDER BY submitted_at DESC;

-- name: DeleteMemory :exec
DELETE FROM memories
WHERE id = $1;

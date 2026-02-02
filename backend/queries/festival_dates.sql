-- name: CreateFestivalDate :one
INSERT INTO festival_dates (
    festival_id, year, start_date, end_date, is_tentative
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetFestivalDatesByFestivalID :many
SELECT * FROM festival_dates
WHERE festival_id = $1
ORDER BY year DESC;

-- name: GetFestivalDateByYear :one
SELECT * FROM festival_dates
WHERE festival_id = $1 AND year = $2;

-- name: ListFestivalDatesByYear :many
SELECT fd.*, f.slug, f.name, f.region, f.heritage_type, f.festival_type, f.summary
FROM festival_dates fd
JOIN festivals f ON f.id = fd.festival_id
WHERE fd.year = $1 AND f.is_published = true
ORDER BY fd.start_date ASC;

-- name: ListUpcomingFestivalDates :many
SELECT fd.*, f.slug, f.name, f.region, f.heritage_type, f.festival_type, f.summary
FROM festival_dates fd
JOIN festivals f ON f.id = fd.festival_id
WHERE f.is_published = true
  AND fd.start_date >= CURRENT_DATE
  AND fd.start_date <= CURRENT_DATE + INTERVAL '30 days'
ORDER BY fd.start_date ASC;

-- name: UpdateFestivalDate :one
UPDATE festival_dates SET
    start_date = $2,
    end_date = $3,
    is_tentative = $4
WHERE id = $1
RETURNING *;

-- name: DeleteFestivalDate :exec
DELETE FROM festival_dates
WHERE id = $1;

-- name: DeleteFestivalDatesByFestivalID :exec
DELETE FROM festival_dates
WHERE festival_id = $1;

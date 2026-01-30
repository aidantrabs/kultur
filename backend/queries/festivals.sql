-- name: ListFestivals :many
SELECT * FROM festivals
WHERE is_published = true
ORDER BY date_2026_start ASC NULLS LAST;

-- name: ListFestivalsByRegion :many
SELECT * FROM festivals
WHERE is_published = true AND region = $1
ORDER BY date_2026_start ASC NULLS LAST;

-- name: ListFestivalsByHeritage :many
SELECT * FROM festivals
WHERE is_published = true AND heritage_type = $1
ORDER BY date_2026_start ASC NULLS LAST;

-- name: ListUpcomingFestivals :many
SELECT * FROM festivals
WHERE is_published = true
  AND date_2026_start >= CURRENT_DATE
  AND date_2026_start <= CURRENT_DATE + INTERVAL '30 days'
ORDER BY date_2026_start ASC;

-- name: GetFestivalBySlug :one
SELECT * FROM festivals
WHERE slug = $1 AND is_published = true;

-- name: GetFestivalByID :one
SELECT * FROM festivals
WHERE id = $1;

-- name: CreateFestival :one
INSERT INTO festivals (
    slug, name, date_type, usual_month, date_2026_start, date_2026_end,
    region, heritage_type, festival_type, summary, story, what_to_expect,
    how_to_participate, practical_info, cover_image_url, gallery_images,
    video_embeds, is_published
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18
) RETURNING *;

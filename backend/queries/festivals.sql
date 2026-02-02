-- name: ListFestivals :many
SELECT * FROM festivals
WHERE is_published = true
ORDER BY name ASC;

-- name: ListFestivalsByRegion :many
SELECT * FROM festivals
WHERE is_published = true AND region = $1
ORDER BY name ASC;

-- name: ListFestivalsByHeritage :many
SELECT * FROM festivals
WHERE is_published = true AND heritage_type = $1
ORDER BY name ASC;

-- name: GetFestivalBySlug :one
SELECT * FROM festivals
WHERE slug = $1 AND is_published = true;

-- name: GetFestivalByID :one
SELECT * FROM festivals
WHERE id = $1;

-- name: CreateFestival :one
INSERT INTO festivals (
    slug, name, date_type, region, heritage_type, festival_type,
    summary, story, what_to_expect, how_to_participate, practical_info,
    cover_image_url, gallery_images, video_embeds, is_published
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
) RETURNING *;

-- name: UpdateFestival :one
UPDATE festivals SET
    slug = $2,
    name = $3,
    date_type = $4,
    region = $5,
    heritage_type = $6,
    festival_type = $7,
    summary = $8,
    story = $9,
    what_to_expect = $10,
    how_to_participate = $11,
    practical_info = $12,
    cover_image_url = $13,
    gallery_images = $14,
    video_embeds = $15,
    is_published = $16
WHERE id = $1
RETURNING *;

-- name: DeleteFestival :exec
DELETE FROM festivals
WHERE id = $1;

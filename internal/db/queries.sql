-- name: SaveBeat :exec
insert into beats ("id", "beatmaker_id", "bpm", "description", "name", "file_path", "image_path", "archive_path", "range_start", "range_end")
values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);

-- name: SaveGenres :copyfrom
insert into beats_genres ("beat_id", "genre_id")
values ($1, $2);

-- name: SaveTags :copyfrom
insert into beats_tags ("beat_id", "tag_id")
values ($1, $2);

-- name: SaveMoods :copyfrom
insert into beats_moods ("beat_id", "mood_id")
values ($1, $2);

-- name: SaveNote :exec
insert into beats_notes ("beat_id", "note_id", "scale")
values ($1, $2, $3);

-- name: GetBeatByID :one
select * from beats where id = $1;

-- name: GetBeatGenreParams :many
select * from genres;

-- name: GetBeatTagParams :many
select * from tags;

-- name: GetBeatMoodParams :many
select * from moods;

-- name: GetBeatNoteParams :many
select * from notes;

-- name: UpdateBeat :one
update beats
set "name" = coalesce(sqlc.narg('name'), "name"),
    "bpm" = coalesce(sqlc.narg('bpm'), "bpm"),
    "description" = coalesce(sqlc.narg('description'), "description"),
    "range_start" = coalesce(sqlc.narg('range_start'), "range_start"),
    "range_end" = coalesce(sqlc.narg('range_end'), "range_end"),
    "is_image_downloaded" = coalesce(sqlc.narg('is_image_downloaded'), "is_image_downloaded"),
    "is_file_downloaded" = coalesce(sqlc.narg('is_file_downloaded'), "is_file_downloaded"),
    "is_archive_downloaded" = coalesce(sqlc.narg('is_archive_downloaded'), "is_archive_downloaded"),
    "updated_at" = now()
where "id" = sqlc.arg('id') and "is_deleted" = false
returning *;

-- name: DeleteBeatGenres :exec
delete from beats_genres where beat_id = $1;

-- name: DeleteBeatTags :exec
delete from beats_tags where beat_id = $1;

-- name: DeleteBeatMoods :exec
delete from beats_moods where beat_id = $1;

-- name: DeleteBeatNotes :exec
delete from beats_notes where beat_id = $1;

-- name: DeleteBeat :exec
update beats
set "is_deleted" = true,
    "updated_at" = now()
where id = $1;

-- name: SaveOwner :exec
insert into beats_owners ("beat_id", "user_id") values ($1, $2);

-- name: GetOwnerByBeatID :one
select * from beats_owners where beat_id = $1;
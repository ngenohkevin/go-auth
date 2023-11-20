-- name: CreateUser :one
INSERT INTO "Users" (id, username, hashed_password, full_name, email)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetUser :one
SELECT * FROM "Users" WHERE id = $1 LIMIT 1;

-- -- name: UpdateUser :one
-- UPDATE "Users"
-- SET hashed_password = COALESCE(sqlc.narg(hashed_password), hashed_password),
--     password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
--     full_name = COALESCE(sqlc.narg(full_name), full_name),
--     email = COALESCE(sqlc.narg(email), email)
-- WHERE id = sqlc.arg(id)
-- RETURNING *;
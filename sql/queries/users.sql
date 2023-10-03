-- name: CreateUser :one 
INSERT INTO users (id, created_at,updated_at,name)
VALUES (s1,s2,s3,s4)
RETURNING *;
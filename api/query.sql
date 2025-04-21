
-- name: GetUserTasks :many
SELECT * FROM tasks WHERE user_id = $1;

-- name: GetTaskNotifications :many
SELECT * FROM task_notifications WHERE task_id = $1;

-- name: GetTaskByID :one
SELECT * FROM tasks WHERE id = $1;

-- name: InsertUser :one
INSERT INTO users (id, name, background)
VALUES ($1, $2, $3)
RETURNING *;

-- name: InsertTask :one
INSERT INTO tasks (id, title_name, background, sticker, is_measured, measurement_unit, sticker_value, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: InsertTaskNotification :one
INSERT INTO task_notifications (id, text, time, task_id, frequency_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFrequenciesByGroup :many
SELECT * FROM frequencies WHERE frequency_group_id = $1;

-- name: GetAllFrequencyGroups :many
SELECT * FROM frequency_groups;

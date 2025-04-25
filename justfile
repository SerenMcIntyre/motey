# Backend

start-api:
  cd api && go run main.go

start-db:
  docker compose up -d

stop-db:
  docker compose down

update-api-orm:
  cd api && sqlc generate

migrate-db:
  cd api && atlas schema apply \
    -u "postgres://postgres:postgres@localhost:5502/postgres?sslmode=disable"  \
    --dev-url "docker://postgres" \
    --to "file://schema.sql"

start-web:
  cd app && bun run dev

start-tauri:
  cd app && bun tauri dev

start-mobile:
  cd app && bun tauri dev android
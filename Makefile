migreatecreate:
	migrate create -ext sql -dir db/migrations -seq video_verification_process_failed

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:root@localhost:5432/postgres?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/postgres?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

run:
	go run application.go
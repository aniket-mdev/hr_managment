migreatecreate:
	migrate create -ext sql -dir db/migrations -seq video_verification_process_failed

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:root@localhost:5432/video_status?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:root@localhost:5432/video_status?sslmode=disable" --verbose down

sqlc:
	sqlc generate

createdb:
	dbmate -u "postgres://chan:0000@localhost:5432/mxuGo?sslmode=disable" create

dropdb:
	dbmate -u "postgres://chan:0000@localhost:5432/mxuGo?sslmode=disable" drop

migrate-up:
	dbmate -u "postgres://chan:0000@localhost:5432/mxuGo?sslmode=disable" up

migrate-down:
	dbmate -u "postgres://chan:0000@localhost:5432/mxuGo?sslmode=disable" down

run:
	go run main.go

watch:
	gow run .

.PHONY: createDb dropDb migrateUp migrateDown run watch 
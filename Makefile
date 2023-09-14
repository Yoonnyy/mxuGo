createDb:
	dbmate -u "postgres://chan:0000@localhost:5432/mxuGo?sslmode=disable" create

dropDb:
	dbmate -u "postgres://chan:0000@localhost:5432/mxuGo?sslmode=disable" drop

migrateUp:
	dbmate -u "postgres://chan:0000@localhost:5432/mxuGo?sslmode=disable" up

migrateDown:
	dbmate -u "postgres://chan:0000@localhost:5432/mxuGo?sslmode=disable" down

.PHONY: createDb dropDb migrateUp migrateDown
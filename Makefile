# import from config file or flag or env
db_source := "file://db/migrations"
db_connection_string := "postgres://postgres:@localhost/l8t1?sslmode=disable"

run:
	go run cmd/cli/main.go -dbcs=$(db_connection_string)
db/migrate:
	migrate -source $(db_source) -database $(db_connection_string) up
db/drop:
	migrate -source $(db_source) -database $(db_connection_string) drop

.PHONY: run db/migrate db/clear db/drop
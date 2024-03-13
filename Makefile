.PHONY: migrate
migrate:
	psql -U postgres -d postgres -h localhost -a -f ./migrations/create_table.sql
	psql -U postgres -d postgres -h localhost -a -f ./migrations/insert_data.sql

serverStart:
	go run ../imdbClone/cmd/server/main.go

clientCHStart:
	go run ../imdbClone/cmd/clientCH/main.go

clientHttpStart:
	go run ../imdbClone/cmd/clientHttp/main.go


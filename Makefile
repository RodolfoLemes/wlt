#!make
include .env

migration-create:
	migrate create -ext sql -dir database/migrations -seq $(name)

migration-up:
	migrate -path "./database/migrations" -database ${DATABASE_URL} up

migration-down:
	migrate -path "./database/migrations" -database ${DATABASE_URL} down 1

migration-force:
	migrate -path "./database/migrations" -database ${DATABASE_URL} force $(version)

doc-update:
	swag init --pd --ot go
	swag fmt

test: ## Run unit-tests with coverage.
	@echo "\033[2m→ Running unit tests...\033[0m"
	go test -count=1 --race -timeout 60s --cover `go list ./... | grep -v integration-tests`

generate-mocks:
	mockery --dir ./domain --case underscore --all
# ==================================================================================== #
# HELPERS
# ==================================================================================== #

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## run: run the cmd/alfred application
.PHONY: run
run:
	go run ./cmd/alfred

## migrations/new name=$1: create a new database migration
.PHONY: migrations/new
migrations/new:
	@echo 'Creating migration file for ${name}...'
	tern new -m ./internal/database/migrations ${name}

## migrations/up: apply all up database migrations
.PHONY: migrations/up
migrations/up: confirm
	@echo 'Running up migrations...'
	tern migrate -m ./internal/database/migrations --conn-string ${ALFRED_DB_DSN}

## migrations/down: apply all down database migrations
.PHONY: migrations/down
migrations/down: confirm
	@echo 'Running down migrations...'
	tern migrate -m ./internal/database/migrations --destination 0 --conn-string ${ALFRED_DB_DSN}

## migrations/target target=$1: go to a specific migration
.PHONY: migrations/target
migrations/target: confirm
	@echo 'Running target migrations...'
	tern migrate -m ./internal/database/migrations --destination ${target} --conn-string ${ALFRED_DB_DSN}

# ==================================================================================== #
# QUALITY CONTROL
# ==================================================================================== #

## tidy: format all .go files, and tidy and vendor module dependencies
.PHONY: tidy
tidy:
	@echo 'Formatting .go files...'
	go fmt ./...
	@echo 'Tidying module dependencies...'
	go mod tidy
	@echo 'Verifying dependencies...'
	go mod verify

## audit: run quality control checks
.PHONY: audit
audit:
	@echo 'Checking module dependencies'
	go mod tidy -diff
	go mod verify
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...
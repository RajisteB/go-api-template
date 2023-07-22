
build:
	@echo "Building..."
	@go build -o bin/depot *.go

run: build
	@echo "Starting application..."
	@./bin/depot

test: 
	@go test -v ./..

migrate-up: 
	@echo "Running db migrations..."
	migrate -database ${DEPOT_PSQL}?sslmode=disable -path sqldb/migrations up

migrate-down:
	@echo "Undoing migration"
	migrate -database ${DEPOT_PSQL}?sslmode=disable -path sqldb/migrations down
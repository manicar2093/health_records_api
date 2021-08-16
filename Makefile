init_dev_env:
	@ echo "Initializing dev db..."
	@ docker-compose -f dev_kit.yml up -d
	@ echo "Creating DB..."
	@ knex migrate:latest
	@ echo "Running seeds..."
	@ knex seed:run
	@ echo "DONE! :D"

db_rollback:
	@ knex migrate:rollback --all

db_fill:
	@ knex migrate:latest
	@ knex seed:run

mocking:
	@ mockery --all

test:
	@ go test ./... -v

coverage:
	@ go test -cover ./...

coverage_html:
	@ go test ./... -coverprofile=coverage.out
	@ go tool cover -html=coverage.out

MIGRATION_PATH=internal/db/migrations
DB_DRIVER=postgres
DB_USER=postgres
DB_PASS=postgres
DB_NAME=go-rest-api-db
DB_HOST=localhost
DB_PORT=5432
DATABASE_URL=${DB_DRIVER}://${DB_USER}:${DB_PASS}@${DB_HOST}/${DB_NAME}?sslmode=disable


upgrade-db:
	migrate -path ${MIGRATION_PATH} -database ${DATABASE_URL} up

downgrade-db:
	migrate -path ${MIGRATION_PATH} -database ${DATABASE_URL} down
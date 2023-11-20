include app.env

network:
	docker network create ${NETWORK}

build:
	docker build -t ${APP_NAME}:latest .

run:
	docker run ${APP_NAME}:latest

#run:
#	docker run --name ${APP_NAME} --network ${NETWORK} -p 8080:8080 -e DB_HOST=${DB_HOST} -e DB_PORT=${DB_PORT} -e DB_NAME=${DB_NAME} -e DB_USER=${DB_USER} -e DB_PASSWORD=${DB_PASSWORD} -d ${APP_NAME}:latest

postgres:
	docker run --name ${DB_DOCKER_CONTAINER} --network ${NETWORK} -p 5432:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:15-alpine

createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}

dropdb:
	docker exec -it ${DB_DOCKER_CONTAINER} dropdb ${DB_NAME}

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate


include app.env

#DB_URL=postgres://root:secret@localhost:5432/go-auth?sslmode=disable
#DB_NAME=go-auth
#USER=root
#PASSWORD=secret
#HOST=localhost
#DB_PORT=5432
#NETWORK=go-auth-network
#DB_DOCKER_CONTAINER=go-auth-db  # docker container name postgres15-alpine
#APP_NAME=go-auth
#DB_SOURCE=postgres://root:secret@localhost:5432/go-auth?sslmode=disable
#DB_DRIVER=postgres
#SERVER_ADDRESS=0.0.0.0:8080
#ACCESS_TOKEN_DURATION=15m
#TOKEN_SYMMETRIC_KEY=38dd0acc8f9929fe9fe8e77cea918247

network:
	docker network create ${NETWORK}

build:
	docker build -t ${APP_NAME}:latest .

run:
	docker run --rm --name ${APP_NAME} --network ${NETWORK} -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@go-auth-db:5432/go-auth?sslmode=disable" ${APP_NAME}:latest

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

test:
	go test -v -cover -short ./...

.PHONY: network build run postgres createdb dropdb migrateup migratedown sqlc test server

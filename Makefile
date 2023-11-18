include app.env

network:
	docker network create ${NETWORK}

build:
	docker build -t ${APP_NAME}:latest .

postgres:
	docker run --name ${DB_DOCKER_CONTAINER} --network ${NETWORK} -p 5432:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:15-alpine

createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}

dropdb:
	docker exec -it ${DB_DOCKER_CONTAINER} dropdb ${DB_NAME}



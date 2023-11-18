include app.env

postgres:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:15-alpine

createdb:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}

dropdb:
	docker exec -it ${DB_DOCKER_CONTAINER} dropdb ${DB_NAME}



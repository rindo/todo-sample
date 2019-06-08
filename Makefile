build:
	docker-compose build

run-api:
	docker-compose run --rm api

run-app:
	docker-compose run --rm app bash

start:
	docker-compose -f docker-compose.yml up

restart:
	docker-compose down
	docker-compose -f docker-compose.yml up

stop:
	docker-compose down

clean:
	rm -rf tmp/postgres/*
	docker-compose down --rmi all

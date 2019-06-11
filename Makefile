build:
	docker-compose build

run-api:
	docker-compose exec api bash

run-app:
	docker-compose exec app bash

start:
	docker-compose -f docker-compose.yml up -d

restart:
	docker-compose down
	docker-compose -f docker-compose.yml up -d

stop:
	docker-compose down

clean:
	rm -rf tmp/postgres/*
	docker-compose down --rmi all

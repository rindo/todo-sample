build:
	docker-compose build

run-api:
	docker-compose -f docker-compose.yml run api

run-app:
	docker-compose -f docker-compose.yml run app bash

start:
	docker-compose -d -f docker-compose.yml up

restart:
	docker-compose down
	docker-compose up -d -f docker-compose.yml

stop:
	docker-compose down

clean:
	docker-compose down --rmi all -v

build:
	docker-compose build

run-api:
	docker-compose exec todo-sample_api_1 bash

run-app:
	docker-compose exec todo-sample_app_1 bash

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

migrate-up:
	docker-compose exec todo-sample_api_1 sql-migrate up

migrate-down:
	docker-compose exec todo-sample_api_1 sql-migrate down

gen-model:
	docker-compose exec todo-sample_api_1 sqlboiler --wipe psql
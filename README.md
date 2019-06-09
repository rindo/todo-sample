# todo sample

# DB

```
# migrate up
docker-compose run api sql-migrate up

# migrate down
docker-compose run api sql-migrate down
```

# Go

```
# ensure
docker-compose run api deq ensure

# run api server
docker-compose run --rm api go run src/api/server/main.go
```

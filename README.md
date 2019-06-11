# todo sample

# DB

```
# migrate up
docker-compose run --rm api sql-migrate up

# migrate down
docker-compose run --rm api sql-migrate down

# create or update model files
docker-compose run --rm sqlboiler --wipe psql
```


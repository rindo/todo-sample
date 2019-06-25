module api

go 1.12

require (
	api/internal/db v0.0.0
	api/internal/models v0.0.0
	github.com/gin-contrib/cors v1.3.0 // indirect
	github.com/lib/pq v1.1.1 // indirect
)

replace (
	api/internal/db v0.0.0 => ./internal/db
	api/internal/models v0.0.0 => ./internal/models
)

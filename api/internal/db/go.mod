module api/internal/db

go 1.12

require (
	github.com/lib/pq v1.1.1
	github.com/volatiletech/sqlboiler v3.4.0+incompatible
	models v0.0.0
)

replace models v0.0.0 => ../models

module service

go 1.12

require (
	db v0.0.0
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.4.0
	github.com/lib/pq v1.1.1 // indirect
	github.com/volatiletech/inflect v0.0.0-20170731032912-e7201282ae8d // indirect
	github.com/volatiletech/sqlboiler v3.4.0+incompatible // indirect
)

replace db => ../internal/db

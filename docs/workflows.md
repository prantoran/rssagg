## Create migrations and handlers for endpoints
- Write migrations in `sql/schema`
- `make migrate`
    - i.e. `cd sql/schema && goose postgres postgres://postgres:secret@localhost:5432/rssagg up && cd ../..`
- Write DB queries in `sql/queries`
- `make sqlc`
    - i.e. `sqlc generate`
- Write the handlers
- Connect the handlers in `main.go`
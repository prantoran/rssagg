```sql
SELECT version()

SELECT * FROM users;
```


## Goose CLI
- For migrations
- `sql/schema`
- Goose runs migrations in order
- `goose postgres postgres://postgres:secret@localhost:5432/rssagg up`
- `goose postgres postgres://postgres:secret@localhost:5432/rssagg down`

## SQLC CLI
- Config: `sqlc.yaml`
- Creates type safe Go cod using SQL statements
- Run `sqlc` from the root of the package
- `sqlc generate`

## Command history

```bash
go mod init github.com/prantoran/rssagg
rm rssagg; go build && ./rssagg
go get github.com/joho/godotenv
go mod tidy
rm rssagg; go build && ./rssagg
go get github.com/go-chi/chi
go mod tidy
go clean -modcache
go get github.com/go-chi/cors

git init -b main
git add .
git commit -m "Setup router boilerplate"
git remote add origin https://github.com
git remote remove origin
git remote add origin https://github.com/prantoran/rssagg
git remote -v
git push -u origin main

go get github.com/lib/pq
``` 


### Install SQL CLI tools
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
sqlc version
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -version
```
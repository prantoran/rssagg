
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

``` 
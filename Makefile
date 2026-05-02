.PHONY: postgres adminer migrate rssagg-net 
# Declare the targets as phony to avoid conflicts with files of the same name
# i.e. do not generate files named postgres, adminer, or migrate when the commands are run

run:
	rm rssagg; go build && ./rssagg

rssagg-net: # default network does not support service/container name resolution
	docker network inspect rssagg-net \ >/dev/null 2>&1 || docker network create rssagg-net

postgres:
	docker run --rm -ti \
	 --network rssagg-net \
	 --name postgres \
	 -e POSTGRES_PASSWORD=secret \
	 -p 5432:5432 \
	 postgres

adminer:
	docker run --rm -ti \
	 --network rssagg-net \
	 --name adminer \
	 -p 8080:8080 \
	 adminer

migrate:
	goose postgres postgres://postgres:secret@localhost:5432/rssagg up

migrate-down:
	goose postgres postgres://postgres:secret@localhost:5432/rssagg down
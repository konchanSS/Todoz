install:
	go get -u github.com/goadesign/goa/...
run:
	go run ./*.go

docker/build:
	docker-compose build

docker/start:
	docker-compose up -d

docker/stop:
	docker-compose down

docker/logs:
	docker-compose logs

docker/clean:
	docker-compose rm

docker/ssh:
	docker exec -it todoz /bin/bash

rundb:
	go run main.go --dbrun

DBNAME:=todoz
ENV:=development

migrate/init:
	mysql -u todoz -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/up:
	sql-migrate up -env=$(ENV)

migrate/down:
	sql-migrate down -env=$(ENV)

migrate/status:
	sql-migrate status -env=$(ENV)

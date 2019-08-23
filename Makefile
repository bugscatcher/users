all: build

build: build_server build_common_migrations

build_server:
	CGO_ENABLED=0 go build -o users

build_common_migrations:
	CGO_ENABLED=0 go build -o migrate_common ./migrations/common/*.go

compose_create_test_db:
	 docker-compose exec postgres createdb -U postgres users_test

compose_drop_test_db:
	 docker-compose exec postgres dropdb -U postgres users_test

compose_recreate_test_db: compose_drop_test_db compose_create_test_db

compose_migrate_test_db: build_common_migrations
	POSTGRESQL_DATABASE=users_test ./migrate_common

test: compose_migrate_test_db
	POSTGRESQL_DATABASE=users_test go test -cover -race -p 10 ./...

generate:
	git submodule foreach git pull origin master
	mkdir -p ./proto/
	cp ./submodules/m-proto/src/main/proto/* ./proto/
	protoc -I=./proto --gofast_out=plugins=grpc:services  `ls ./proto/`

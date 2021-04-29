all: clean check-quality build test modtidy

ENV_FILE="covid-ed-service.env"

BINARY="./bin/covid-ed-service"

ifeq ("$(environment)", "dev")
	DB_NAME=covid_ed_service_dev
	ENV_FILE="covid-ed.env"
endif

modtidy:
	GO111MODULE=on go mod tidy -v

clean: modtidy
	rm -rf ./out/covid-ed-service

check-quality: setup lint vet imports

golangci:
	golangci-lint run -v --deadline 3m0s

setup:
	GO111MODULE=off go get -v golang.org/x/tools/cmd/goimports
	GO111MODULE=off go get -v golang.org/x/lint/golint
	GO111MODULE=off go get -v github.com/golangci/golangci-lint/cmd/golangci-lint
	GO111MODULE=off	go get -v -d github.com/golang-migrate/migrate/cmd/migrate
	cd "${GOPATH}src/github.com/golang-migrate/migrate/cmd/migrate" && git checkout v4.9.1

	mkdir -p ./out
	echo ${ENV_FILE}


test: setup db.migrate
	go test -race ./...

only_test:
	go test -race ./...

imports:
	goimports -l -w .

lint:
	golint  ./... | grep -iEv 'exported.*should.*comment' || true

vet:
	go vet ./...

build:
	go build -o ${BINARY} ./cmd/server
	GOOS=linux GOARCH=amd64 go build -ldflags "-w" -o  ./bin/covid-ed-server ./cmd/server

db.migrate:
	migrate -verbose -path migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" up

db.rollback:
	migrate -verbose -path migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" down 1 

db.rollback_all:
	echo Y | migrate -verbose -path migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" down

db.drop:
	dropdb ${DB_NAME} --if-exists

db.create:
	createdb ${DB_NAME}

db.seed:
	psql "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -a -f scripts/seed_data.sql

db.clear:
	psql "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}" -a -f scripts/clear.sql

db.login:
	psql "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSL_MODE}"

LOCAL_BIN=$(CURDIR)/bin

.PHONY: build start test

.EXPORT_ALL_VARIABLES:
GO111MODULE = on
APP_STAGE = local

include bin-deps.mk

lint: $(GOLANGCI_BIN)
	$(GOLANGCI_BIN) run ./...

build:
	CGO_ENABLED=0 go build -ldflags="-s -w" -o ./bin/identity cmd/identity-server/main.go

# db should be already up before starting server
run: build
	./bin/identity --port 8080 --db.username="identity" --db.password="identitypass" --db.dbname="identity"  --db.host="0.0.0.0" --db.port=5432

test: gotest gofmt govet

gotest:
	go test ./... -v -cover

gofmt:
	go fmt ./...

govet:
	go vet ./...

gogenerate:
	go generate ./... -v

gen-server:
	swagger generate server -A identity -f ./api/spec/server.yaml --server-package=./internal/server/restapi --model-package=./pkg/models/rest --exclude-main --principal rest.Principal --with-context

new-migration:
ifndef name
	 @echo "need pass \"migration name\" variable";
else
	/usr/local/bin/migrate create -dir migrations -ext sql ${name}
endif

up:
	docker-compose -f docker-compose.yaml down
	docker-compose -f docker-compose.yaml up -d

migrate-local:
	migrate -path ./migrations -database postgresql://identity:identitypass@localhost:5432/identity?sslmode=disable up
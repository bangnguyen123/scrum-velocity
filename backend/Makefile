build:
	@docker-compose build

up:
	@docker-compose up -d

log-backend:
	@docker logs backend -f --tail 200

prisma-migrate:
	@docker exec backend go run github.com/steebchen/prisma-client-go db push

prisma-generate:
	@docker exec backend go run github.com/steebchen/prisma-client-go generate

build-local:
	go build main.go

install-local:
	go mod download

run-local:
	@air

prisma-generate-local:
	go run github.com/steebchen/prisma-client-go generate

prisma-migrate-local:
	go run github.com/steebchen/prisma-client-go db push

unit-test:
	go test ./test/...

lint:
	golangci-lint run
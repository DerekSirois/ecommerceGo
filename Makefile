run:
	@go run cmd/web/*

up_build: build_ecommerce
	docker-compose down
	docker-compose up --build -d
	docker image prune -f
	rm ecommerceGo

up:
	docker-compose up -d

down:
	docker-compose down

build_ecommerce:
	env GOOS=linux CGO_ENABLED=0 go build -o ecommerceGo ./cmd/web

test:
	@go test ./internal/...
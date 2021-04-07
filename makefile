build: format
	docker-compose build --no-cache

build_linux: format
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o build/quizen -v ./cmd/quizen

down:
	docker-compose down

format:
	go fmt ./...

start:
	docker-compose start

stop:
	docker-compose stop

test:
	go test -v ./...

up:
	docker-compose up -d app

.PHONY: \
	build \
	build_linux \
	down \
	format \
	start \
	stop \
	test \
	up
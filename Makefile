.PHONY: build run docker-build docker-run

build:
	go build -o eulabs cmd/server/main.go

run: build
	./eulabs

docker-build:
	docker build -t eulabs .

docker-run:
	docker run -e DB_TYPE=$(DB_TYPE) -p 8080:8080 eulabs
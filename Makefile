BINARY=build

build:
	go build -o bin/engine main.go

run:
	 go run main.go

docker:
	docker build -t ghcr.io/jiramot/go-profile-resource . --no-cache

push:
	docker push ghcr.io/jiramot/go-profile-resource
	
.PHONY: run build
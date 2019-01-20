
.PHONY: all image

all: app

image: Dockerfile
	docker build . -f Dockerfile -t go-api-image

app: main.go
	go build -o app main.go

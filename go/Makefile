
.PHONY: all image deploy

all: deploy

image: Dockerfile
	docker build . -f Dockerfile -t go-api-image

app: main.go
	go build -o app main.go

deploy: image
	docker run --rm -p 8081:8081 go-api-image ./app

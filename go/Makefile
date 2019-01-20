
.PHONY: all image deploy

TAG := go-api-image

all: deploy

image: Dockerfile
	docker build . -f Dockerfile -t $(TAG)

app: app.go
	go build -o app app.go

deploy: image
	docker run --rm -p 8081:8080 $(TAG)

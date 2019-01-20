
.PHONY: all image deploy

all: deploy

image: Dockerfile
	docker build . -f Dockerfile -t go-api-image

app: app.go
	go build -o app app.go

deploy: image
	docker run --rm -p 8081:8080 go-api-image ./app

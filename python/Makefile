
.PHONY: all image deploy

TAG := python-api-image

all: deploy

image:
	docker build . -f Dockerfile -t $(TAG)

deploy: image
	docker run --rm -p 8080:8080 $(TAG)

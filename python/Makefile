
.PHONY: all image deploy

all: deploy

image:
	docker build . -f Dockerfile -t flask-image

deploy: image
	docker run --rm -p 8080:8080 flask-image ./app.py

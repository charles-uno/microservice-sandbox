
.PHONY: all image deploy

all: deploy

image:
	docker build . -f Dockerfile -t flask-image

deploy: image
	docker run --rm -p 5000:5000 flask-image ./app.py

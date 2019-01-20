
# Python-Based REST API with Flask

Creates a microservice within a Docker container. To fire up the microservice locally, run

```
make
```

This will pull the dependencies and API source into a Docker image, then run the API within that image, publishing the container's port 8080 to `localhost:8080`. To see it in action, point your browser to `localhost:8080/factors/N` for some positive integer `N` (or use `curl`). It should spit out a (JSON-formatted) list of that number's factors. For example, `localhost:8080/factors/357` shows:

```
[
  3,
  7,
  17
]
```

This service is built using Python's `connexion` module, which is itself a Swagger-processing layer on top of Flask. Note that this is a proof of concept, and not meant as a template for deployment. [Additional tools](http://flask.pocoo.org/docs/1.0/tutorial/deploy/#) are needed to properly deploy Flask code.

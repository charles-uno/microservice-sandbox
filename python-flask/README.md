
# Python-Based REST API with Flask

Creates a microservice within a Docker container. To fire up the microservice locally, run

```
make
```

Then navigate to `localhost:5000/api/factors/N` for some positive integer `N`. For example, a `GET` operation on `localhost:5000/api/factors/357` returns:

```
[
  3,
  7,
  17
]
```

This service is built using Python's `connexion` module, which is itself a Swagger-processing layer on top of Flask.

Note that this is a proof of concept, and not meant as a template for deployment. [Additional tools](http://flask.pocoo.org/docs/1.0/tutorial/deploy/#) are needed to properly deploy Flask code. 

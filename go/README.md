
# Go-Based REST API with net/http

Creates a microservice within a Docker container. To fire up the microservice locally, run

```
make
```

This will pull the dependencies and API source into a Docker image, then run the API within that image, publishing the container's port 8080 to `localhost:8081`. To see it in action, point your browser to `localhost:8081/factors/N` for some positive integer `N` (or use `curl`). It should spit out a (JSON-formatted) list of that number's factors. For example, `localhost:8081/factors/357` shows:

```
[
  3,
  7,
  17
]
```

This service uses only tools in the Go standard library. There are some third-party tools that could clean up the parameter input, but this is enough for our purposes.

**TODO:** Check this against the Swagger spec.

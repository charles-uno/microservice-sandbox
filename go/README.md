
# Go-Based REST API with net/http

Creates a microservice within a Docker container. To fire up the microservice locally, run

```
make
```

Then point your browser to `localhost:8081/factors/N` for some positive integer `N` (or use `curl`). For example, a `GET` operation on `localhost:8081/factors/357` returns:

```
[
  3,
  7,
  17
]
```

This service uses only tools in the Go standard library. There are some third-party tools that could clean up the parameter input, but this is enough for our purposes.

**TODO:** Check this against the Swagger spec.

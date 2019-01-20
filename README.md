
# Microservice Sandbox

Playing around with making REST APIs, with Swagger specs, and sticking them in Docker containers.

Basically the same API is implemented in both Go and Python. You give it a positive integer, and it returns a (JSON-formatted) list of that integer's prime factors.

In both cases, the build process looks about like this:

- Pull dependencies into a Docker image.
- Build/install the API source into the image.
- Run the container, publishing the relevant port.

Both containers map their service to (container) port 8080. The Python-based API is published to `localhost:8080` and the Go-based API is published to `localhost:8081` to avoid collisions if they are running concurrently.

Next up is to figure out the Swagger magic that allows these APIs to be plugged into Kubernetes or whatever.

Note that debug messages and error responses are a function of the implementation frameworks, and may not match verbatim between Python and Go.

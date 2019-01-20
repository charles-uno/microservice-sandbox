
# Microservice Sandbox

Playing around with making REST APIs, with Swagger specs, and sticking them in Docker containers.

Basically the same API is implemented in both Go and Python. You give it a positive integer, and it returns a (JSON-formatted) list of that integer's prime factors.

In both cases, the build process looks about like this:

- Pull dependencies into a Docker image.
- Build/install the API source into the image.
- Run the container, publishing the relevant port.

For Python I use port 5000. For Go, I use 8081. These are just conventions, and readily changed. Additionally, Docker's `--publish` function allows arbitrary mapping between container ports and host ports. In principle, all containerized APIs could use (container) port 5000 and they could be dynamically mapped to host ports as needed. 

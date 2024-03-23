# DCUpdate

DCUpdate is a simple tool to update your Docker containers. It is designed to be used in a CI/CD pipeline, and responds to a HTTP request, that will check the version of a container and update it if necessary.

DCUpdate is not intended for production use, but rather for development and testing purposes.  If you want to do this in production, you should use a more robust solution like Kubernetes.

In my use case, I have 2 API's each in separate containers, 2 front end applications. Exposed to the internet using Caddy. They are all serviced from a small VM running docker compose. DCUpdate is also in the docker compose file, and exposed through caddy on an update URL. These are not mission critical applications, and this is a cheap way and easy to update them.

When a new version of a container needs to be updated. The CI/CD pipeline will send a HTTP request to the update URL, and DCUpdate will check the version of the container and update it, then restart it. 

Another project that do something similar is Watchtower however, it is not as flexible as DCUpdate, DCUpdate has a update end point and will except version parameters for the container





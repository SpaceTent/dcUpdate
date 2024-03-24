# dcUpdate

### dcUpdate is a simple tool designed to be called from your CI/CD pipline to update Docker containers. 

This is not intended for production use, but rather for development and testing purposes. If you want to do this in production, you should use a more robust solution like Kubernetes, if your not changing tags on the containers, you can use Watchtower.

In my specific use case, I have 4 applications in containers Exposed to the internet using Caddy. They are all serviced from a small VM running docker compose. When a new version of a container needs to be updated. The CI/CD pipeline will send a HTTP request to the update URL, and dcUpdate will check the version of the container, download the new image if needed, rewrite the config file then restart it. This is hacky, but it works for my use case. I did orginally go down the route for re-creating all the containers using the Docker API however I needed to replicate the config loading functionality in the docker-compose file, which is not trivial. Everything bar that, is just about there.  I'll come back to that later. 

1) Starts a small HTTP Server
2) Listens for a GET request
3) When an update request is received, it will 
    a) Get the List of Containers running and check this requests for the container name
    b) if this requires a new image, it will pull the new image. 
        1) If this requires Auth, this creds are obtained from an ENV variable

    c) Update the Docker Compose config file with the new image tag
    d) Restart everything










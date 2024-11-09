# Needed commands
```docker build -t <TAGNAME> .```
Builds the docker image. A Dockerfile is required in the current directory.

```docker run <TAGNAME>```
Creates a container from the image with <TAGNAME> and starts it.

```docker image ls -a```
Shows all images that are stored on the system where the docker daemon is running.

```docker container ls -a```
Shows all containers on the system.

```docker container stop <CONTAINER_ID>```
Stops the container with the id <CONTAINER_ID>, if running.

```docker compose up```
Starts a docker compose and attaches compose's console to terminal.

```docker compose up -d```
Starts a docker compose, but does not attach the compose's console to the terminal.

```docker compose down```
Stops and removes all containers and networks from the docker compose.
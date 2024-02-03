# Docker image creation

The following commands are used to create the docker images for the frontend. The docker image will be created and pushed to the GitHub Container Registry.
It will have the tag `v1`.

```bash
docker login ghcr.io
cd ./frontend
docker build -t ghcr.io/noleu/event_calendar_frontend:v1 .
docker push ghcr.io/noleu/event_calendar_frontend:v1
``` 

In general is that not the intended way, as a github workflow exists and triggered by a push on main. 
The workflow can be found in [here](./.github/workflows/publish-ghcr.yml).
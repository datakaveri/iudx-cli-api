# City Intelligence Layer APIs

## Setup

### Dev

- Run **air** command to start the server in dev mode

### Prod

- Build the docker image and deploy using below commands

```bash
cd setup
docker compose build
docker compose up -d
```

Note:

- If you are getting error while building the image due to module download error its because of ipv6
- [Enable ipv6](https://docs.docker.com/config/daemon/ipv6/#use-ipv6-for-the-default-bridge-network) in docker and restart docker to fix the issue.

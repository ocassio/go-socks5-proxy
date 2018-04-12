# Build
The following command will perform a build of a static binary for Linux.
The result of this build can be used by a scratch Docker image. This reduces container size drastically.

```bash
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .
```

# Docker image usage

You can start Docker container with the following command (make sure to replace `<USER>` and `<PASSWORD>` placeholders with your own credentials).

```bash
docker run -d --name socks5-proxy -p 1080:1080 -e USER=<USER> -e PASSWORD=<PASSWORD> ocassio/go-socks5-proxy
```

# Special thanks

- The original idea. Actually this project is just a very lightweight variation of the serj's one:
   https://github.com/serjs/socks5-server
- SOCKS 5 server implementation for Go:
   https://github.com/armon/go-socks5
- Article about building minimal Go containers with Docker by Nick Gauthier:
   https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/
